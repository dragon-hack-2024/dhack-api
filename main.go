package main

import (
	"dhack-api/api"
	"dhack-api/config"
	"log"
)

func main() {

	// Load configuration settings.
	config, err := config.New(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Create a server and setup routes.
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Failed to create a server: ", err)
	}

	// Start a server.
	if err := server.Start(config.Address); err != nil {
		log.Fatal("Failed to start a server: ", err)
	}
}
