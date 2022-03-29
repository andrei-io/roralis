package main

import (
	"backend/roralis/config"
	"backend/roralis/routes"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Launches the app.
// Looks for configuration and errors out if not found.
// Starts up dependecy injection container.
// Starts the web server.
func main() {

	if err := config.Load(); err != nil {
		fmt.Printf("Failed to read config:\n     %v", err)
		os.Exit(1)
	}

	appConfig, err := config.BootstrapServices()
	if err != nil {
		fmt.Printf("Failed to assemble services:\n 	%v", err)
		os.Exit(1)
	}

	app := config.BuildApp(appConfig)

	// By default, Gin uses the PORT enviroment variables

	// Mounts the routes
	routes.MountRoutes(app, appConfig)

	fmt.Printf("Server running on localhost:%v\n", viper.GetString("PORT"))
	err = app.Run(":" + viper.GetString("PORT"))
	if err != nil {
		fmt.Printf("Failed to start Gin server:\n    %v", err)
		os.Exit(1)
	}

}
