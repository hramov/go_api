package api

import (
	ioc "api/src/core/container"
	"api/src/core/logger"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func initRestApi(prefix string) {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	engine = router

	api := router.Group(prefix)
	ioc.Put("Router", api)
}

func createRestServer() {
	if engine == nil {
		logger.Error("Gin engine not initialized")
		return
	}
	if err := engine.Run(":" + os.Getenv("APP_PORT")); err != nil {
		logger.Error("Cannot start REST server: " + err.Error())
	}
}
