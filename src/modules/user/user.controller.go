package user

import (
	ioc "api/src/core/container"
	"api/src/utils"

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
	utils.SendResponse(200, "", result, c)
}
