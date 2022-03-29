package config

import (
	"backend/roralis/auth"
	"backend/roralis/category"
	postController "backend/roralis/controllers/post"
	userController "backend/roralis/controllers/user"
	"backend/roralis/domain/repo/email"
	"backend/roralis/domain/repo/post"
	"backend/roralis/otc"
	"backend/roralis/region"

	"backend/roralis/domain/repo/user"
	"backend/roralis/infrastructure"
	"backend/roralis/middleware"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// nolint: govet
// Services struct for all the dependencies
type Services struct {
	TokenKey string
	DB       *gorm.DB

	JWTSecret  *auth.JWTSecret
	JWTService auth.JWTService

	CategoryRepo       category.CategoryRepo
	CategoryController category.CategoryController

	RegionRepo       region.RegionRepo
	RegionController region.RegionController

	PostRepo       post.PostRepo
	PostController postController.PostController

	UserRepo       user.UserRepo
	UserController userController.UserController

	EmailRepo email.EmailRepo
	OTCRepo   otc.OTCRepo

	AuthService middleware.AuthService
}

// Set up all the services.
// Will error out.
func BootstrapServices() (*Services, error) {
	var (
		err    error
		config Services
	)
	config.DB, err = infrastructure.NewDB(viper.GetString("DB_URL"))
	if err != nil {
		return nil, err
	}
	config.TokenKey = "token"

	config.CategoryRepo = category.NewCategoryRepo(config.DB)
	config.CategoryController = category.NewCategoryController(config.CategoryRepo)

	config.RegionRepo = region.NewRegionRepo(config.DB)
	config.RegionController = region.NewRegionController(config.RegionRepo)

	config.PostRepo = post.NewPostRepo(config.DB)
	config.PostController = postController.NewPostController(config.PostRepo, config.TokenKey)

	config.EmailRepo = email.NewEmailRepo(viper.GetString("SENDGRID_KEY"))
	config.OTCRepo = otc.NewOTCRepo(config.DB)

	config.JWTSecret, err = loadRSAKeys(viper.GetString("JWT_PRIVATE"), viper.GetString("JWT_PUBLIC"))
	if err != nil {
		return nil, err
	}
	config.JWTService = auth.NewJWTService(config.JWTSecret)

	config.UserRepo = user.NewUserRepo(config.DB)
	config.UserController = userController.NewUserController(
		config.UserRepo,
		config.EmailRepo,
		config.OTCRepo,
		config.JWTService,
		config.TokenKey,
	)

	config.AuthService = middleware.NewAuthService(config.JWTService, config.TokenKey)

	return &config, nil

}
