package guards

import (
	auth_dto "api/src/modules/auth/dto"
	"api/src/modules/auth/jwt"
	"api/src/modules/logger"
	"api/src/modules/user"
	"api/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"golang.org/x/crypto/bcrypt"
)

func LocalGuard(c *gin.Context) {

	body := utils.GetBody[auth_dto.LoginDto](c)

	var userService *user.UserService
	if err := container.NamedResolve(&userService, "UserService"); err != nil {
		logger.Error("Cannot resolve User service")
	}

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
