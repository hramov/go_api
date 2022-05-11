package api

import (
	ioc "api/src/core/container"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) Init(prefix string) {
	s.router = gin.New()
	gin.SetMode(gin.ReleaseMode)

	s.router.Use(gin.Recovery())
	s.router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	s.router.Use(cors.New(config))

	api := s.router.Group(prefix)
	ioc.Put("Router", api)
}

func (s *Server) Start() {
	s.router.Run(":" + os.Getenv("APP_PORT"))
}
