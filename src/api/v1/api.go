package api

import (
	"api/src/core/logger"
	"api/src/modules/auth"
	"api/src/modules/user"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func (a *Server) Start() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	api := router.Group("/api/v1")

	auth.Init(api)
	user.Init(api)

	logger.Info("Initialized all router groups")
	router.Run(":3000")
}
