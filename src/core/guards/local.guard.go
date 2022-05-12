package guards

import (
	ioc "api/src/core/container"
	"api/src/core/jwt"
	auth_dto "api/src/modules/auth/dto"
	user_entity "api/src/modules/user/entity"
	"api/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LocalGuard(c *gin.Context) {

	body, err := utils.GetBody[auth_dto.LoginDto](c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	userService := ioc.Pick[*user_entity.UserRepository]("UserRepository")

	user, err := userService.FindByEmail(body.Email)

	if err == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": fmt.Errorf("Cannot find user").Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := jwt.CreateToken(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("access_token", token)
	c.Next()
}
