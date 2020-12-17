package main

import (
	"os"
	"os/signal"

	"github.com/lyticaa/lyticaa-app/internal/web"
)

func main() {
	a := app.NewApp()
	a.Start()

	defer a.Data.Cache.Quit()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.Stop()
}
