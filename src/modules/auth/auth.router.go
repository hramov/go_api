package auth

import (
	ioc "api/src/core/container"
	"api/src/core/guards"
	"api/src/core/pipes"
	auth_dto "api/src/modules/auth/dto"

	"github.com/gin-gonic/gin"
)

func InitRouter(controller *AuthController) {
	router := ioc.Pick[*gin.RouterGroup]("Router")

	auth := router.Group("/auth")
	{
		auth.GET("/info",
			guards.JwtAuthGuard([]string{}),
			controller.UserInfo)

		auth.POST("/login",
			pipes.ValidationPipe[auth_dto.LoginDto](),
			guards.LocalGuard,
			controller.Login)
	}
}
