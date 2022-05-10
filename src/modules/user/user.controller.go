package user

import (
	"net/http"

	user_port "api/src/modules/user/port"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service user_port.UserServicePort
}

func (uc *UserController) Find(c *gin.Context) {
	result := uc.service.Find()
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
