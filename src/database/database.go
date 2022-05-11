package database

import (
	ioc "api/src/core/container"
	"api/src/core/logger"
	user_entity "api/src/modules/user/entity"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(migrate bool) {
	PostgresDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(PostgresDSN), &gorm.Config{})
	if err != nil {
		logger.Error("Cannot connect to database with DSN: " + PostgresDSN)
	}

	if migrate {
		migrateModels(db)
	}

	ioc.Put("postgres", db)
}

func migrateModels(db *gorm.DB) {
	if err := db.AutoMigrate(&user_entity.User{}); err != nil {
		logger.Error("Cannot migrate User")
	}
}
