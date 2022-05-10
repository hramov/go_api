package main

import (
	"api/src/api/v1"
	ioc "api/src/core/container"
	"api/src/core/logger"
	"api/src/database"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file")
	}

	database.InitPostgres()

	ioc.Init()
	logger.Info("IoC started")

	api := api.Server{}
	go api.Start()

	logger.Info("Server gorutine started")

	for {

	}
}
