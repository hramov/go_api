package auth

import (
	auth_dto "api/src/modules/auth/dto"
	"api/src/modules/logger"
	user_entity "api/src/modules/user/entity"
	"api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
)

type AuthController struct{}

func (ac *AuthController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, &gin.H{
		"message": "Pong",
	})
}

func (ac *AuthController) Login(c *gin.Context) {

	var authService *AuthService
	if err := container.NamedResolve(&authService, "AuthService"); err != nil {
		logger.Error("Cannot resolve AuthService")
	}

	dto, err := utils.GetBody[auth_dto.LoginDto](c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := authService.Login(dto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	utils.SendResponse(http.StatusOK, "Successfilly logged in", token, c)
	return
}

func (ac *AuthController) UserInfo(c *gin.Context) {
	user, exists := c.Get("user")

	if user == nil || !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "User not found",
		})
		return
	}
	utils.SendResponse(http.StatusOK, "Successfilly logged in", user.(*user_entity.User), c)
}
