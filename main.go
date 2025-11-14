package main

import (
	"log"
	"os"
	"tpIRSO/app"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}

	application := app.NewApp()
	log.Println("Starting server on port:", port)
	application.Start(port)
}
