package auth

import (
	"api/src/core/guards"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {

	controller := &AuthController{}

	auth := router.Group("/auth")
	{
		auth.GET("/ping", controller.Ping)
		auth.GET("/info", guards.JwtAuthGuard([]string{}), controller.UserInfo)
		auth.POST("/login", guards.LocalGuard, controller.Login)
	}
}
