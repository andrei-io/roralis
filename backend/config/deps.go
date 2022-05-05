package config

import (
	"backend/roralis/core/auth"
	"backend/roralis/core/category"
	"backend/roralis/core/jwt"
	"backend/roralis/core/post"
	posthttp "backend/roralis/core/post/post_http"
	"backend/roralis/core/region"
	"backend/roralis/core/user"

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

	JWTSecret  *jwt.JWTSecret
	JWTService jwt.JWTService

	CategoryRepo       category.CategoryRepo
	CategoryController category.CategoryController

	RegionRepo       region.RegionRepo
	RegionController region.RegionController

	PostRepo       post.PostRepo
	PostController posthttp.PostController

	UserRepo       user.UserRepo
	UserController user.UserController

	AuthController auth.AuthController

	AuthMiddleware middleware.AuthService
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
	config.PostController = posthttp.NewPostController(config.PostRepo, config.TokenKey)

	config.JWTSecret, err = loadRSAKeys(viper.GetString("JWT_PRIVATE"), viper.GetString("JWT_PUBLIC"))
	if err != nil {
		return nil, err
	}
	config.JWTService = jwt.NewJWTService(config.JWTSecret)

	config.UserRepo = user.NewUserRepo(config.DB)
	config.UserController = user.NewUserController(config.UserRepo)

	config.AuthController = auth.NewAuthController(
		config.UserRepo,
		config.JWTService,
		config.TokenKey,
	)

	config.AuthMiddleware = middleware.NewAuthService(config.JWTService, config.TokenKey)

	return &config, nil

}
