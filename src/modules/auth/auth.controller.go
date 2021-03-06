package auth

import (
	ioc "api/src/core/container"
	user_entity "api/src/modules/user/entity"
	"api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *AuthService
}

func createController() *AuthController {
	return &AuthController{
		Service: ioc.Pick[*AuthService]("AuthService"),
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	token, exists := c.Get("access_token")

	if !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Token not exists",
		})
		return
	}

	utils.SendResponse(http.StatusOK, "Successfilly logged in", token.(string), c)
}

func (ac *AuthController) UserInfo(c *gin.Context) {
	user, exists := c.Get("user")

	if user == nil || !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "User not found",
		})
		return
	}
	utils.SendResponse(http.StatusOK, "", user.(*user_entity.User), c)
}
