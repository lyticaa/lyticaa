package main

import (
	"github.com/lyticaa/lyticaa-app/internal/worker"
)

func main() {
	w := worker.NewWorker()
	w.Start()
}
