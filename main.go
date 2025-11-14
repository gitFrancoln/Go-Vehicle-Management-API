package main

import (
	"tpIRSO/internal/config"
	app "tpIRSO/src/appp"
)

func main() {
	config.LoadConfig()
	app := app.NewApp()

	app.Start(config.Env.Port)
}
