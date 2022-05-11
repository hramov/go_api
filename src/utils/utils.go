package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetReqResFromContext(c *gin.Context) (*http.Request, *http.Response) {
	return c.Request, c.Request.Response
}

func GetTokenFromRequest(req *http.Request) (string, error) {
	auth := req.Header.Get("authorization")
	if auth != "" {
		cred := strings.Split(auth, " ")
		if len(cred) > 1 && cred[0] == "Bearer" {
			if cred[1] != "" {
				return cred[1], nil
			}
			return "", fmt.Errorf("No token")
		}
		return "", fmt.Errorf("Wrong auth header format")
	}
	return "", fmt.Errorf("No auth header")
}

func GetTokenFromContext(c *gin.Context) (string, error) {
	req, _ := GetReqResFromContext(c)
	return GetTokenFromRequest(req)
}

func Exists[T comparable](array []T, value T) bool {
	if len(array) == 0 {
		return true
	}

	for _, val := range array {
		if (val) == value {
			return true
		}
	}
	return false
}

func SendResponse[T comparable](status int, message string, data T, c *gin.Context) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func SendError(status int, err error, c *gin.Context) {
	c.AbortWithStatusJSON(status, err)
}

func GetBody[T comparable](c *gin.Context) T {
	var data T
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return data
	}

	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		return data
	}
	return data
}

func CheckErrorForHttp(err error, status int, c *gin.Context) {
	if err != nil {
		// c.AbortWithStatusJSON(status, gin.H{
		// 	"error": err.Error(),
		// })
		c.AbortWithError(status, err)
		return
	}
}

func GetPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

/**
This function inits all modules.
Modules provide their services in IoC container.
Modules can be resolved by inner function Get<ModuleName>Module()
Modules services can be resolved via ioc.Pick[T](name string) T
*/

type Initable interface {
	Init()
}

func InitModules(modules []Initable) {
	for i := 0; i < len(modules); i++ {
		temp := modules[i]
		temp.Init()
	}
}
