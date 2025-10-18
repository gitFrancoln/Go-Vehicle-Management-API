package main

import (
	"tpIRSO/app"
	"tpIRSO/internal/config"
)

func main() {
	config.LoadConfig()
	app := app.NewApp()

	app.Start(config.Env.Port)
}
