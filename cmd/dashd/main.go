package main

import (
	"os"
	"os/signal"

	"gitlab.com/getlytica/lytica-app/internal/core/app"
)

func main() {
	a := app.NewApp()
	a.Start()

	defer a.Cache.Quit()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.Stop()
}
