package jwt

import (
	user_entity "api/src/modules/user/entity"
	"fmt"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}

func CreateToken(user *user_entity.User) (string, error) {
	atClaims := Claims{}
	atClaims.Id = strconv.FormatUint(uint64(user.ID), 10)
	atClaims.Email = user.Email
	atClaims.Role = user.Role
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

func GetPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
