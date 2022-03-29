package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Creates Gin server and mounts routes
func BuildApp(c *Services) *gin.Engine {
	gin.SetMode(viper.GetString("GIN_MODE"))

	// Blank gin app.
	// Not using gin.Default because in the future the app will be using custom logging and recovery methods
	app := gin.New()
	app.Use(gin.Recovery())
	// TODO: use a logging library here
	app.Use(gin.Logger())
	app.Use(gin.ErrorLogger())

	return app
}
