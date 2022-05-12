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
	InitRouter(um.controller)
}

func GetUserModule() *UserModule {
	if userModule == nil {
		logger.Error("UserModule not initialized")
	}
	return userModule
}
