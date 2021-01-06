package data

import (
	"sync"

	"github.com/lyticaa/lyticaa/internal/data/reports/ingest"
)

type Worker interface {
	Start() error
}

func (d *Data) Start() {
	wg := &sync.WaitGroup{}

	for _, w := range d.workers() {
		wg.Add(1)

		go func(w Worker, wg *sync.WaitGroup) {
			defer wg.Done()

			if err := w.Start(); err != nil {
				d.Monitoring.Logger.Error().Err(err).Msg("failed to start worker")
			}
		}(w, wg)
	}

	wg.Wait()
}

func (d *Data) workers() []Worker {
	var workers []Worker
	workers = append(workers, ingest.NewIngest(d.Monitoring.NewRelic, d.Database.PG))

	return workers
}
