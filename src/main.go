package main

import (
	"api/src/api/v1"
	"api/src/core/logger"
	"api/src/database"
	"api/src/modules/auth"
	"api/src/modules/user"
	"api/src/utils"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		logger.Error("Error loading .env file")
	} // Load .env file with configs

	database.InitPostgres(true) // Establish connect to database and migrate models if needed
	logger.Info("Successfilly connected to database")

	api := api.Server{}
	api.Init("/api/v1")

	utils.InitModules([]utils.Initable{
		&auth.AuthModule{},
		&user.UserModule{},
	}) // Init registered modules, provide their services in IoC container

	logger.Info("Server gorutine started")

	api.Start()
}
