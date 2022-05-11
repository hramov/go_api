package auth

import (
	ioc "api/src/core/container"
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
	am.Controller = &AuthController{}
	am.Service = &AuthService{}
	authModule = am

	ioc.Put("AuthService", am.Service)
}
