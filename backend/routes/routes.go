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
	"country/controllers/category"
	"country/controllers/posts"
	"country/controllers/region"
	"country/controllers/user"
	_ "country/doc" // for swagger responses

	"github.com/gin-gonic/gin"
)

// Mounts the routes
func MountRoutes(app *gin.Engine) {
	v1 := app.Group("/api/v1")

	// swagger:route GET /users/:id user getUser
	//
	// Get one user
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GetOneUserResponse
	v1.GET("/users/:id", user.ReadOne)

	// swagger:route POST /users/signup user signup
	//
	// Sign Up
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignUpSucces
	v1.POST("/users/signup", user.SignUp)

	// swagger:route GET /users/aboutme user aboutme
	//
	// Sign Up
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     AboutMeSucces
	v1.GET("/users/aboutme", IsLoggedIn, user.AboutMe)

	// swagger:route POST /users/confirm/:id user confirm
	//
	// Validate email
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignInSucces
	v1.POST("/users/confirm/:id", user.ValidateEmail)

	// swagger:route GET /users/validate user confirm
	//
	// Resend email
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GenericResponse
	v1.GET("/users/resend/:id", user.ResendValidationEmail)

	// swagger:route GET /users/validate user confirm
	//
	// Resend email
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignInSucces
	v1.POST("/users/signin", user.SignIn)

	// swagger:route GET /categories/ category getCategory
	//
	// Get all categories
	//
	//     Responses:
	//       200:     GetAllCategoriesResponse
	//       default: GenericResponse
	v1.GET("/categories", category.ReadAll)

	// swagger:route GET /categories/:id category getOneCategory
	//
	// Get category by id
	//
	//     Responses:
	//       200:     GetOneCategoriesResponse
	//       default: GenericResponse
	v1.GET("/categories/:id", category.ReadOne)

	// swagger:route GET /regions/ region getRegion
	//
	// Get all regions
	//
	//     Responses:
	//       200:     GetAllRegionsResponse
	//       default: GenericResponse
	v1.GET("/regions", region.ReadAll)

	// swagger:route GET /regions/:id region getCategory
	//
	// Get region by id
	//
	//     Responses:
	//       200:     GetOneRegionsResponse
	//       default: GenericResponse
	v1.GET("/regions/:id", region.ReadOne)

	// swagger:route GET /regions/:id region getCategory
	//
	// Get region by id
	//
	//     Responses:
	//       200:     GetOnePostResponse
	//       default: GenericResponse
	v1.GET("/posts/:id", posts.ReadOne)

}
