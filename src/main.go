package main

import (
	"api/src/api/v1"
	"api/src/core/logger"
	"api/src/database"
	"api/src/modules"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
	}

	database.InitPostgres(true)
	logger.Info("IoC started")

	modules.InitModules()

	api := api.Server{}
	go api.Start()

	logger.Info("Server gorutine started")

	for {

	}
}
