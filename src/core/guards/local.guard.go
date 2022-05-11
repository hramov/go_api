package guards

import (
	ioc "api/src/core/container"
	"api/src/core/jwt"
	auth_dto "api/src/modules/auth/dto"
	"api/src/modules/user"
	"api/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LocalGuard(c *gin.Context) {

	body := utils.GetBody[auth_dto.LoginDto](c)

	userService := ioc.Pick[*user.UserService]("UserService")

	user := userService.FindByEmail(body.Email)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": fmt.Errorf("Cannot find user"),
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
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
