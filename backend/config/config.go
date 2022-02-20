package config

import (
	"errors"
	"os"
	"strings"

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
	viper.SetDefault("SENDGRID_KEY", "")
	viper.SetDefault("EMAIL_NAME", "")
	viper.SetDefault("EMAIL_FROM", "")
	viper.SetDefault("ENV", "")
	viper.SetDefault("REDIS_ADDRESS", "")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("JWT_PUBLIC", "")
	viper.SetDefault("JWT_PRIVATE", "")

	_, err = os.Stat(".env")
	if !errors.Is(err, os.ErrNotExist) {
		viper.SetConfigFile(".env")
		viper.SetConfigType("props")
		err = viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	// Override config parameters from environment variables if specified
	for _, key := range viper.AllKeys() {
		if err := viper.BindEnv(key); err != nil {
			return err
		}
	}

	// Some .env readers dont read things properly

	key := viper.GetString("JWT_PRIVATE")

	key = strings.ReplaceAll(key, `\n`, "\n")

	viper.Set("JWT_PRIVATE", key)

	key = viper.GetString("JWT_PUBLIC")

	key = strings.ReplaceAll(key, `\n`, "\n")

	viper.Set("JWT_PUBLIC", key)

	return nil
}
