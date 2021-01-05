package main

import (
	"os"
	"os/signal"

	"github.com/lyticaa/lyticaa-app/internal/app"
)

func main() {
	a := app.NewApp()
	a.APIHandlers()

	a.Start(false)

	defer a.Database.Memcache.Quit()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.Stop()
}
