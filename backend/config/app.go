package config

import (
	"country/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Starts Gin server and mounts routes
func StartApp() (*gin.Engine, error) {
	gin.SetMode(viper.GetString("GIN_MODE"))

	app := gin.New()
	app.Use(gin.Recovery())
	// TODO: use a logging library here
	app.Use(gin.Logger())
	app.Use(gin.ErrorLogger())

	routes.MountRoutes(app)

	return app, nil
}
