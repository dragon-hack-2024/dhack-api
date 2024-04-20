package main

import (
	"dhack-api/api"
	"dhack-api/config"
	"dhack-api/db"
	"log"
)

func main() {

	// Load configuration settings.
	config, err := config.New(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	store, err := db.Connect(config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}
	defer store.Close()

	// Create a server and setup routes.
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Failed to create a server: ", err)
	}

	// Start a server.
	if err := server.Start(config.Address); err != nil {
		log.Fatal("Failed to start a server: ", err)
	}
}
