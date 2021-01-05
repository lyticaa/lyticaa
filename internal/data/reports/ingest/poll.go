package ingest

import (
	"os"

	"github.com/bufferapp/sqs-worker-go/worker"
)

func (i *Ingest) Start() error {
	i.Logger.Info().Msgf("listening for messages on %v....", os.Getenv("AWS_SQS_QUEUE"))

	w, err := worker.NewService(os.Getenv("AWS_SQS_QUEUE"))
	if err != nil {
		i.Logger.Error().Err(err).Msg("unable to create new worker service")
		return err
	}

	w.Start(worker.HandlerFunc(i.Parse))

	return nil
}
