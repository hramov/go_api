package user

import (
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

	controller := &UserController{
		service: &UserService{
			Repo: &user_entity.UserRepository{
				Db: db,
			},
		},
	}

	user := router.Group("/user")
	{
		user.GET("/", controller.Find)
	}
}
