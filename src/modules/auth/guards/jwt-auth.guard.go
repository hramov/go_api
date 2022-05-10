package guards

import (
	"api/src/modules/auth/jwt"
	"api/src/modules/logger"
	"api/src/modules/user"
	user_entity "api/src/modules/user/entity"
	"api/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

func JwtAuthGuard(c *gin.Context, roles []string) {

	var db *gorm.DB
	if err := container.NamedResolve(&db, "postgres"); err != nil {
		logger.Error("Cannot resolve db")
	}

	req, _ := utils.GetReqResFromContext(c)
	token, err := utils.GetTokenFromRequest(req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := jwt.TokenValid(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := data["jti"].(string)

	userService := &user.UserService{
		Repo: &user_entity.UserRepository{
			Db: db,
		},
	}

	user := userService.FindBy("id", id)

	if user != nil {
		c.Set("user", user)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": fmt.Errorf("User not found"),
	})
	return
}
