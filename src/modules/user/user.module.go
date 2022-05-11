package user

import (
	"api/src/core/logger"
)

type UserModule struct {
	Controller *UserController
	Service    *UserService
}

var userModule *UserModule

func (um *UserModule) Init() {
	um.Service = createService()
	um.Controller = createController()
	userModule = um
	InitRouter()
}

func GetUserModule() *UserModule {
	if userModule == nil {
		logger.Error("UserModule not initialized")
	}
	return userModule
}

func getController() *UserController {
	return userModule.Controller
}

func getService() *UserService {
	return userModule.Service
}
