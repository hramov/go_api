package main

/**

Name: Golang REST API server (inspired by NestJS)
Version: 0.0.1
Author: Khramov Sergey <trykhramov@gmail.com> | github.com/hramov

*/

import (
	"api/src/api/v1"
	"api/src/core/logger"
	"api/src/database"
	"api/src/modules/auth"
	"api/src/modules/user"
	"api/src/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		logger.Error("Error loading .env file")
	} // Load .env file with configs

	database.InitPostgres(true) // Establish connect to database and migrate models if needed
	logger.Info("Successfilly connected to database")

	api.Init("/api/v1")

	utils.InitModules([]utils.Initable{
		&auth.AuthModule{},
		&user.UserModule{},
	}) // Init registered modules, provide their services in IoC container

	api.Start()

	logger.Info("Application successfully started at port:\n\t\t\t\t - " + os.Getenv("APP_PORT") + " (REST)\n\t\t\t\t - " + os.Getenv("GRPC_PORT") + " (GRPC)")
	for {
	}
}
