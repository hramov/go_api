package user

import (
	ioc "api/src/core/container"
	"api/src/core/guards"

	"github.com/gin-gonic/gin"
)

func InitRouter(controller *UserController) {
	router := ioc.Pick[*gin.RouterGroup]("Router")

	user := router.Group("/user")
	{
		user.GET("/", guards.JwtAuthGuard([]string{"admin"}), controller.Find)
	}
}
