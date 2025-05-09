package main

import (
	"restapp/internal/app"
)

func main() {
	app := app.NewApp("../../config/config.yaml")
	app.Run()
}
