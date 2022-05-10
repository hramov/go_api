package auth

import (
	auth_dto "api/src/modules/auth/dto"
	"api/src/modules/auth/jwt"
	user_port "api/src/modules/user/port"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo user_port.UserRepoPort
}

func (as *AuthService) Login(dto auth_dto.LoginDto) (string, error) {

	user := as.repo.FindByEmail(dto.Email)
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
