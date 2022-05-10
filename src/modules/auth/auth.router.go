package auth

import (
	"api/src/modules/auth/guards"
	"api/src/modules/logger"
	user_entity "api/src/modules/user/entity"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

func Init(router *gin.RouterGroup) {

	var db *gorm.DB
	if err := container.NamedResolve(&db, "postgres"); err != nil {
		logger.Error("Cannot resolve db")
	}

	controller := &AuthController{
		service: &AuthService{
			repo: &user_entity.UserRepository{
				Db: db,
			},
		},
	}

	auth := router.Group("/auth")
	{
		auth.GET("/ping", controller.Ping)
		auth.GET("/info", func(c *gin.Context) { guards.JwtAuthGuard(c, []string{}) }, controller.UserInfo)
		auth.POST("/login", controller.Login)
	}
}
