package user

import (
	ioc "api/src/core/container"
	"api/src/core/guards"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := ioc.Pick[*gin.RouterGroup]("Router")
	controller := getController()

	user := router.Group("/user")
	{
		user.GET("/", guards.JwtAuthGuard([]string{"admin"}), controller.Find)
	}
}
