package main

import (
	"gitlab.com/getlytica/dashboard/internal/worker/app"
)

func main() {
	a := app.NewApp()
	a.Start()
}
