package main

import (
	"gitlab.com/getlytica/lytica/internal/worker/app"
)

func main() {
	a := app.NewApp()
	a.Start()
}
