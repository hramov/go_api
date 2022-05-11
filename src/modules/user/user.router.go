package user

import (
	ioc "api/src/core/container"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := ioc.Pick[*gin.RouterGroup]("Router")
	controller := getController()

	user := router.Group("/user")
	{
		user.GET("/", controller.Find)
	}
}
