package config

import (
	"boiler-go/handler"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	// app mode
	appMode := os.Getenv("APP_MODE")
	switch appMode {
	case "STAGING":
		gin.SetMode(gin.TestMode)
	case "PRODUCTION":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// init gin routes
	route := gin.Default()

	// global var
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// versioning prefix
	api := route.Group("/api/v1")

	// routes list
	api.GET("/dummy", handler.HelloDummy)

	// routes 404
	route.NoRoute(handler.PageNotFound)

	return route
}
