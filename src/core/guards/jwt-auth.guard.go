package guards

import (
	"api/src/core/jwt"
	"api/src/core/logger"
	"api/src/modules/user"
	"api/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

func JwtAuthGuard(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userService *user.UserService
		if err := container.NamedResolve(&userService, "UserService"); err != nil {
			logger.Error("Cannot resolve UserService")
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
