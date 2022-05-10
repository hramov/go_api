package auth

import (
	"api/src/modules/auth/guards"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {

	controller := &AuthController{}

	auth := router.Group("/auth")
	{
		auth.GET("/ping", controller.Ping)
		auth.GET("/info", func(c *gin.Context) { guards.JwtAuthGuard(c, []string{}) }, controller.UserInfo)
		auth.POST("/login", controller.Login)
	}
}
