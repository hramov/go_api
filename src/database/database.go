package database

import (
	"api/src/core/logger"
	user_entity "api/src/modules/user/entity"
	"fmt"
	"os"

	"github.com/golobby/container/v3"
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

	logger.Info("Successfilly connected to database")

	if migrate {
		migrateModels(db)
	}

	err = container.NamedSingleton("postgres", func() *gorm.DB {
		return db
	})
	if err != nil {
		logger.Error("Cannot use IoC")
	}
}

func migrateModels(db *gorm.DB) {
	if err := db.AutoMigrate(&user_entity.User{}); err != nil {
		logger.Error("Cannot migrate User")
	}
}
