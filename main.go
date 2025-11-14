package main

import "tpIRSO/app"

func main() {
	application := app.NewApp()
	application.Start("10000") // Render expone el puerto dinámico, pero 10000 es el que usa internamente
}
