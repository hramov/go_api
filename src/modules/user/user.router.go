package user

import (
	"api/src/core/logger"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

func Init(router *gin.RouterGroup) {

	var db *gorm.DB
	if err := container.NamedResolve(&db, "postgres"); err != nil {
		logger.Error("Cannot resolve db")
	}

	controller := &UserController{}

	user := router.Group("/user")
	{
		user.GET("/", controller.Find)
	}
}
