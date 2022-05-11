package user

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {

	controller := &UserController{}

	user := router.Group("/user")
	{
		user.GET("/", controller.Find)
	}
}
