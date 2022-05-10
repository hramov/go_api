package user

import (
	"api/src/core/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

type UserController struct{}

func (uc *UserController) Find(c *gin.Context) {

	var userService *UserService
	if err := container.NamedResolve(&userService, "UserService"); err != nil {
		logger.Error("Cannot resolve UserService")
	}

	result := userService.Find()
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
