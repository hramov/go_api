package user

import (
	"api/src/core/logger"
)

type UserModule struct {
	controller *UserController
	Service    *UserService
}

var userModule *UserModule

func (um *UserModule) Init() {
	um.Service = createService()
	um.controller = createController()
	userModule = um
	initRouter(um.controller)
	logger.Info("User module successfilly initialized")
}

func GetUserModule() *UserModule {
	if userModule == nil {
		logger.Error("UserModule not initialized")
	}
	return userModule
}
