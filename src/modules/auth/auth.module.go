package auth

import (
	"api/src/core/logger"
)

type AuthModule struct {
	Controller *AuthController
	Service    *AuthService
}

var authModule *AuthModule

func GetAuthModule() *AuthModule {
	if authModule == nil {
		logger.Error("AuthModule not initialized")
	}
	return authModule
}

func (am *AuthModule) Init() {
	am.Service = createService()
	am.Controller = createController()
	authModule = am
	InitRouter()
}

func getController() *AuthController {
	return authModule.Controller
}

func getService() *AuthService {
	return authModule.Service
}
