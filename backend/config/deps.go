package config

import (
	categoryController "backend/roralis/controllers/category"
	"backend/roralis/domain/entity"
	"backend/roralis/domain/repo/category"
	"backend/roralis/domain/repo/email"
	"backend/roralis/domain/repo/otc"
	"backend/roralis/domain/repo/post"
	"backend/roralis/domain/repo/region"
	"backend/roralis/domain/repo/user"
	"backend/roralis/domain/services/jwt"
	"backend/roralis/infrastructure"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// // nolint: govet
// Config struct for all the dependencies
type Config struct {
	TokenKey           string
	DB                 *gorm.DB
	CategoryRepo       category.CategoryRepo
	CategoryController categoryController.CategoryController

	UserRepo   user.UserRepo
	RegionRepo region.RegionRepo
	EmailRepo  email.EmailRepo
	OTCRepo    otc.OTCRepo
	PostRepo   post.PostRepo
	JWTSecret  entity.JWTSecret
	JWTService jwt.JWTService
}

// Set up all the services.
// Will error out.
func BootstrapServices() (*Config, error) {
	var (
		err    error
		config Config
	)
	config.TokenKey = "token"

	config.DB, err = infrastructure.NewDB(viper.GetString("DB_URL"))
	if err != nil {
		return nil, err
	}
	config.CategoryRepo = category.NewCategoryRepo(config.DB)
	config.CategoryController = categoryController.NewCategoryController(config.CategoryRepo)

	return &config, nil

}
