package user

import (
	ioc "api/src/core/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *UserService
}

var controllerInstance *UserController

func createController() *UserController {
	if controllerInstance == nil {
		controllerInstance = &UserController{
			Service: ioc.Pick[*UserService]("UserService"),
		}
	}
	return controllerInstance
}

func (uc *UserController) Find(c *gin.Context) {
	result := uc.Service.Find()
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
