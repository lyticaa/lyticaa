package app

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gitlab.com/getlytica/lytica/internal/worker/app/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (a *App) parseMessage(msg *sqs.Message) error {
	body := []byte(*msg.Body)
	var r types.Response

	err := json.Unmarshal(body, &r)
	if err != nil {
		a.Logger.Error().Err(err)
		return nil
	}

	for _, record := range r.Records {
		a.Logger.Info().Msgf("processing %v....", record.S3.Object.Key)
		go a.processFile(record.S3.Object.Key)
	}

	return nil
}

func (a *App) processFile(file string) {
	cfg := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		a.Logger.Fatal().Err(err)
	}

	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_UPLOAD_BUCKET")),
		Key:    aws.String(file),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		a.Logger.Error().Err(err)
	}

	validType := types.ValidMime(*result.ContentType)
	username := a.userFromKey(file)

	if validType {

		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			a.Logger.Error().Err(err)
		}

		rows := a.mapCsv(file, bytes.NewBuffer(body))
		a.Logger.Info().Str("user", username).Msgf("total rows to process: %v", len(rows))
	} else {
		a.Logger.Info().Str("user", username).Msgf("invalid content type: %v", *result.ContentType)
	}
}

func (a *App) mapCsv(file string, reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	var rows []map[string]string
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if len(record) > 0 {
			shouldSkip := types.ShouldIgnore(record[0])
			if shouldSkip {
				a.Logger.Info().Str("user", a.userFromKey(file)).Msgf("skipping record: %v", record[0])
				continue
			}
		}

		if err != nil {
			a.Logger.Fatal().Err(err)
		}

		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}

	return rows
}

func (a *App) userFromKey(key string) string {
	parts := strings.Split(key, "/")
	return parts[0]
}
