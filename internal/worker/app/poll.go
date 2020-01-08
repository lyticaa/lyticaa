package app

import (
	"os"

	"github.com/bufferapp/sqs-worker-go/worker"
)

func (a *App) Start() {
	a.Logger.Info().Msg("starting....")

	w, err := worker.NewService(os.Getenv("AWS_SQS_QUEUE"))
	if err != nil {
		a.Logger.Error().Err(err).Msg("unable to create new worker service")
	}

	w.Start(worker.HandlerFunc(a.parseMessage))
}
