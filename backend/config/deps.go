package config

import (
	categoryController "backend/roralis/controllers/category"
	postController "backend/roralis/controllers/post"
	regionController "backend/roralis/controllers/region"
	userController "backend/roralis/controllers/user"
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
	TokenKey string
	DB       *gorm.DB

	JWTSecret  *entity.JWTSecret
	JWTService jwt.JWTService

	CategoryRepo       category.CategoryRepo
	CategoryController categoryController.CategoryController

	RegionRepo       region.RegionRepo
	RegionController regionController.RegionController

	PostRepo       post.PostRepo
	PostController postController.PostController

	UserRepo       user.UserRepo
	UserController userController.UserController

	EmailRepo email.EmailRepo
	OTCRepo   otc.OTCRepo
}

// Set up all the services.
// Will error out.
func BootstrapServices() (*Config, error) {
	var (
		err    error
		config Config
	)
	config.DB, err = infrastructure.NewDB(viper.GetString("DB_URL"))
	if err != nil {
		return nil, err
	}
	config.TokenKey = "token"

	config.CategoryRepo = category.NewCategoryRepo(config.DB)
	config.CategoryController = categoryController.NewCategoryController(config.CategoryRepo)

	config.RegionRepo = region.NewRegionRepo(config.DB)
	config.RegionController = regionController.NewRegionController(config.RegionRepo)

	config.PostRepo = post.NewPostRepo(config.DB)
	config.PostController = postController.NewPostController(config.PostRepo)

	config.EmailRepo = email.NewEmailRepo(viper.GetString("SENDGRID_KEY"))
	config.OTCRepo = otc.NewOTCRepo(config.DB)

	config.JWTSecret, err = loadRSAKeys(viper.GetString("JWT_PRIVATE"), viper.GetString("JWT_PUBLIC"))
	if err != nil {
		return nil, err
	}
	config.JWTService = jwt.NewJWTService(config.JWTSecret)

	config.UserRepo = user.NewUserRepo(config.DB)
	config.UserController = userController.NewUserController(
		config.UserRepo,
		config.EmailRepo,
		config.OTCRepo,
		config.JWTService,
	)

	return &config, nil

}