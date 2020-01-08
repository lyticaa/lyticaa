package main

import (
	"os"
	"os/signal"

	"gitlab.com/getlytica/dashboard/internal/dashboard/app"
)

func main() {
	a := app.NewApp()
	a.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.Stop()
}
