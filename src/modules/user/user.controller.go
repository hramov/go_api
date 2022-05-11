package user

import (
	ioc "api/src/core/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) Find(c *gin.Context) {

	userService := ioc.Pick[*UserService]("UserService")

	result := userService.Find()
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
