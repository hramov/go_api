package auth

import (
	"api/src/core/logger"
)

type AuthModule struct {
	controller *AuthController
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
	am.controller = createController()
	authModule = am
	InitRouter(am.controller)
}
