package main

import (
	_ "github.com/lib/pq"
	"ideahive/backend/cmd/app"
	"log"
)

func main() {
	newApp, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	log.Fatal(newApp.Serve())
}
