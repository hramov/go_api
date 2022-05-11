package auth

import ioc "api/src/core/container"

type AuthService struct{}

func createService() *AuthService {
	service := &AuthService{}
	ioc.Put("AuthService", service)
	return service
}
