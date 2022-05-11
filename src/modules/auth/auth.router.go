package auth

import (
	ioc "api/src/core/container"
	"api/src/core/guards"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := ioc.Pick[*gin.RouterGroup]("Router")
	controller := getController()

	auth := router.Group("/auth")
	{
		auth.GET("/info", guards.JwtAuthGuard([]string{}), controller.UserInfo)
		auth.POST("/login", guards.LocalGuard, controller.Login)
	}
}
