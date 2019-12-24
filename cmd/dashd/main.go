package main

import (
	"os"
	"os/signal"

	"gitlab.com/sellernomics/dashboard/internal/dash"
)

func main() {
	d := dash.NewDash()
	d.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	d.Stop()
}
