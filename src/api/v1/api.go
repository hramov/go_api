package api

import (
	ioc "api/src/core/container"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (a *Server) Init() {
	a.router = gin.New()
	gin.SetMode(gin.ReleaseMode)

	a.router.Use(gin.Recovery())
	a.router.Use(gin.Logger())

	api := a.router.Group("/api/v1")
	ioc.Put("Router", api)
}

func (a *Server) Start() {
	a.router.Run(":" + os.Getenv("APP_PORT"))
}
