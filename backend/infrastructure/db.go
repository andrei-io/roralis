// Contains everything related physical infrastructure
package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Returns a new DB
// Gets its configuration from viper
func NewDB(dbUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(dbUrl),
		&gorm.Config{},
	)

	return db, err
}
