package main

import (
	"gitlab.com/getlytica/lytica-app/internal/worker/app"
)

func main() {
	a := app.NewApp()
	a.Start()
}
