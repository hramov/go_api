package user

import (
	ioc "api/src/core/container"
	"api/src/core/logger"
	user_entity "api/src/modules/user/entity"

	"gorm.io/gorm"
)

type UserModule struct {
	Controller *UserController
	Service    *UserService
}

var userModule *UserModule

func GetUserModule() *UserModule {
	if userModule == nil {
		logger.Error("UserModule not initialized")
	}
	return userModule
}

func (um *UserModule) Init() {

	db := ioc.Pick[*gorm.DB]("postgres")
	um.Controller = &UserController{}
	um.Service = &UserService{
		&user_entity.UserRepository{
			Db: db,
		},
	}
	userModule = um

	ioc.Put("UserService", um.Service)
}
