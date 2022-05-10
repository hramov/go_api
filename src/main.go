package main

import (
	"api/src/api/v1"
	"api/src/database"
	"api/src/modules/logger"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		logger.Error("Error loading .env file")
	}

	database.InitPostgres()
	// Request handlers
	api := api.Server{}
	go api.Start()

	logger.Info("Server gorutine started")

	for {

	}
}
