package auth

import (
	auth_dto "api/src/modules/auth/dto"
	"api/src/modules/auth/jwt"
	"api/src/modules/logger"
	"api/src/modules/user"
	"fmt"

	"github.com/golobby/container/v3"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (as *AuthService) Login(dto auth_dto.LoginDto) (string, error) {

	var userService *user.UserService
	if err := container.NamedResolve(&userService, "UserService"); err != nil {
		logger.Error("Cannot resolve User service")
	}

	user := userService.FindByEmail(dto.Email)
	if user == nil {
		return "", fmt.Errorf("Cannot find user")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return "", err
	}

	token, err := jwt.CreateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}
