package config

import (
	"github.com/spf13/viper"
)

// Reads config from "base.env", ".env", and system enviroment.
// Prefers env variables
func Load() error {
	var err error

	// Set some sane defaults
	// TODO: set more
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("GIN_MODE", "release")
	viper.SetDefault("DB_URL", "host=localhost user=admin password=admin dbname=asdf sslmode=disable")

	viper.SetConfigFile(".env")
	viper.SetConfigType("props")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Override config parameters from environment variables if specified
	for _, key := range viper.AllKeys() {
		if err := viper.BindEnv(key); err != nil {
			return err
		}
	}
	return nil
}
