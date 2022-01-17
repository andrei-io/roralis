package infrastructure

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dbUrl := viper.GetString("DB_URL")
	db, err := gorm.Open(
		postgres.Open(dbUrl),
		&gorm.Config{},
	)

	if err != nil {
		// Dont wanna have to return an error so just panicking is fine for now
		panic("failed to connect to db")
	}

	return db
}
