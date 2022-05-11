package guards

import (
	ioc "api/src/core/container"
	"api/src/core/jwt"
	user_entity "api/src/modules/user/entity"
	"api/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthGuard(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userService := ioc.Pick[*user_entity.UserRepository]("UserRepository")

		token, err := utils.GetTokenFromContext(c)

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

		user := userService.FindBy("id", id)

		if user != nil {
			if utils.Exists(roles, user.Role) {
				c.Set("user", user)
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": fmt.Errorf("User must have different role").Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": fmt.Errorf("User not found").Error(),
		})
		return
	}
}
