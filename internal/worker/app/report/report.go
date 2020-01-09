package report

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gitlab.com/getlytica/lytica/internal/worker/app/report/types"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	customTransaction = "CustomTransaction"
	unknown           = int64(1)
)

type Report struct {
	Logger zerolog.Logger
	Db     *sqlx.DB
}

func NewReport(db *sqlx.DB) *Report {
	return &Report{
		Logger: log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
		Db:     db,
	}
}

func (r *Report) Run(file string) {
	result := r.getS3Object(file)
	validType := types.ValidMime(*result.ContentType)
	username := r.userFromKey(file)

	if validType {
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			r.Logger.Error().Err(err)
			return
		}

		var inputType string
		if strings.Contains(file, "CustomTransaction") {
			inputType = "CustomTransaction"
		}

		r.processReport(inputType, username, body)
	} else {
		r.Logger.Info().Str("user", username).Msgf("invalid content type: %v", *result.ContentType)
	}
}

func (r *Report) processReport(file, username string, body []byte) {
	rows := r.mapCsv(file, bytes.NewBuffer(body))
	r.Logger.Info().Str("user", username).Msgf("total rows to process: %v", len(rows))

	switch file {
	case customTransaction:
		txns := r.formatTransactions(rows, username)
		for _, txn := range txns {
			err := r.saveTransaction(txn)
			if err != nil {
				r.Logger.Error().Err(err)
			}
		}
	}
}

func (r *Report) mapCsv(file string, reader io.Reader) []map[string]string {
	rr := csv.NewReader(reader)
	var rows []map[string]string
	var header []string
	for {
		record, err := rr.Read()
		if err == io.EOF {
			break
		}

		if len(record) > 0 {
			shouldSkip := types.ShouldIgnore(record[0])
			if shouldSkip {
				r.Logger.Info().Str("user", r.userFromKey(file)).Msgf("skipping record: %v", record[0])
				continue
			}
		}

		if err != nil {
			r.Logger.Fatal().Err(err)
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

func (r *Report) userFromKey(key string) string {
	parts := strings.Split(key, "/")
	return parts[0]
}
