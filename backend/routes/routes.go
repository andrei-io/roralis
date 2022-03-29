// Application description
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package routes

import (
	"backend/roralis/config"
	"backend/roralis/controllers/user"
	_ "backend/roralis/doc" // for swagger responses

	"github.com/gin-gonic/gin"
)

// Mounts the routes
func MountRoutes(app *gin.Engine, c *config.Config) {
	v1 := app.Group("/api/v1")

	// swagger:route GET /api/v1/users/:id user getUser
	//
	// Get one user
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GetOneUserResponse
	v1.GET("/users/:id", user.ReadOne)

	// swagger:route POST /api/v1/users/signup user signup
	//
	// Sign Up
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignUpSucces
	v1.POST("/users/signup", user.SignUp)

	// swagger:route GET /api/v1/users/aboutme user aboutme
	//
	// Sign Up
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     AboutMeSucces
	v1.GET("/users/aboutme", IsLoggedIn, user.AboutMe)

	// swagger:route POST /api/v1/users/confirm/:id user confirm
	//
	// Validate email
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignInSucces
	v1.POST("/users/confirm/:id", user.ValidateEmail)

	// swagger:route GET /api/v1/users/validate user resend
	//
	// Resend email
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GenericResponse
	v1.GET("/users/resend/:id", user.ResendValidationEmail)

	// swagger:route GET /api/v1/users/validate user signin
	//
	// Resend email
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignInSucces
	v1.POST("/users/signin", user.SignIn)

	// swagger:route GET /api/v1/categories/ category getCategory
	//
	// Get all categories
	//
	//     Responses:
	//       200:     GetAllCategoriesResponse
	//       default: GenericResponse
	v1.GET("/categories", c.CategoryController.ReadAll)

	// swagger:route GET /api/v1/categories/:id category getOneCategory
	//
	// Get category by id
	//
	//     Responses:
	//       200:     GetOneCategoriesResponse
	//       default: GenericResponse
	v1.GET("/categories/:id", c.CategoryController.ReadOne)

	// swagger:route GET /api/v1/regions/ region getRegion
	//
	// Get all regions
	//
	//     Responses:
	//       200:     GetAllRegionsResponse
	//       default: GenericResponse
	v1.GET("/regions", c.RegionController.ReadAll)

	// swagger:route GET /api/v1/regions/:id region getCategory
	//
	// Get region by id
	//
	//     Responses:
	//       200:     GetOneRegionsResponse
	//       default: GenericResponse
	v1.GET("/regions/:id", c.RegionController.ReadOne)

	// swagger:route GET /api/v1/posts/:id posts getPost
	//
	// Get region by id
	//
	//     Responses:
	//       200:     GetOnePostResponse
	//       default: GenericResponse
	v1.GET("/posts/:id", c.PostController.ReadOne)

	// swagger:route GET /api/v1/posts/?offset=&limit=20 posts getAllPosts
	//
	// Get post by id
	//
	//     Responses:
	//       200:     GetAllPostResponse
	//       default: GenericResponse
	v1.GET("/posts", c.PostController.ReadAll)

	// swagger:route POST /api/v1/posts/ posts createPost
	//
	// Create post
	//
	//     Responses:
	//       200:     GetOnePostResponse
	//       default: GenericResponse
	v1.POST("/posts", IsLoggedIn, c.PostController.Create)

}
