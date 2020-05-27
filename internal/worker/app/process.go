package app

import (
	"encoding/json"

	"gitlab.com/getlytica/lytica-app/internal/worker/app/report"
	"gitlab.com/getlytica/lytica-app/internal/worker/app/types"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func (a *App) parseMessage(msg *sqs.Message) error {
	body := []byte(*msg.Body)
	var rr types.Response

	err := json.Unmarshal(body, &rr)
	if err != nil {
		a.Logger.Error().Err(err)
		return nil
	}

	r := report.NewReport(a.Db)

	for _, record := range rr.Records {
		a.Logger.Info().Msgf("processing %v....", record.S3.Object.Key)
		go r.Run(record.S3.Object.Key)
	}

	return nil
}
