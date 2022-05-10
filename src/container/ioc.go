package ioc

import (
	"api/src/modules/auth"
	"api/src/modules/logger"
	"api/src/modules/user"
	user_entity "api/src/modules/user/entity"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

func Pick[T comparable](name string, where T) {
	if err := container.NamedResolve(&where, name); err != nil {
		logger.Error("Cannot resolve " + name)
	}
}

func Init() {
	initServices()
}

func initServices() {

	var db *gorm.DB
	if err := container.NamedResolve(&db, "postgres"); err != nil {
		logger.Error("Cannot resolve db")
	}

	err := container.NamedSingleton("UserService", func() *user.UserService {
		return &user.UserService{
			Repo: &user_entity.UserRepository{
				Db: db,
			},
		}
	})

	err = container.NamedSingleton("AuthService", func() *auth.AuthService {
		return &auth.AuthService{}
	})

	if err != nil {
		logger.Error(err.Error())
	}

}
