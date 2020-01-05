package main

import (
	"os"
	"os/signal"

	"gitlab.com/getlytica/dashboard/internal/core"
)

func main() {
	c := core.NewCore()
	c.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	c.Stop()
}
