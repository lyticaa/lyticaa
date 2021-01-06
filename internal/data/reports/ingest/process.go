package ingest

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/lyticaa/lyticaa/internal/data/reports/ingest/amazon"
	"github.com/lyticaa/lyticaa/internal/data/reports/ingest/types"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func (i *Ingest) Parse(msg *sqs.Message) error {
	body := []byte(*msg.Body)
	var rr types.Response

	err := json.Unmarshal(body, &rr)
	if err != nil {
		i.Logger.Error().Err(err)
		return nil
	}

	amz := amazon.NewAmazon(i.Db)

	for _, record := range rr.Records {
		username, filename := i.fileParts(record.S3.Object.Key)

		if !amz.Reports.ShouldProcess(filename) {
			i.Logger.Info().Msgf("skipping %v as not a valid Amazon report....", filename)
		} else {
			i.Logger.Info().Msgf("processing %v for %v....", filename, username)
			wg := &sync.WaitGroup{}
			wg.Add(1)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()

				amz.Reports.Run(username, filename)
			}(wg)

			wg.Wait()
		}
	}

	return nil
}

func (i *Ingest) fileParts(filename string) (string, string) {
	return strings.Split(filename, "/")[0], filename
}
