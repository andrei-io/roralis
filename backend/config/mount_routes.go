// Package classification Roralis API.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package config

import (
	_ "backend/roralis/doc" // for swagger responses

	"github.com/gin-gonic/gin"
)

// Mounts the routes
func mountRoutes(app *gin.Engine, c *Services) {
	v1 := app.Group("/api/v1")

	// swagger:route GET /api/v1/users/:id user getUser
	//
	// Get one user
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GetOneUserResponse
	v1.GET("/users/:id", c.UserController.ReadOne)

	// swagger:route POST /api/v1/users/signup user signup
	//
	// Sign Up
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignUpSucces
	v1.POST("/users/signup", c.AuthController.SignUp)

	// swagger:route GET /api/v1/users/aboutme user aboutme
	//
	// About me
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     AboutMeSucces
	v1.GET("/users/aboutme", c.AuthMiddleware.IsLoggedIn, c.AuthController.AboutMe)

	// swagger:route POST /api/v1/users/signin user signin
	//
	// Sign In
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     SignInSucces
	v1.POST("/users/signin", c.AuthController.SignIn)

	// swagger:route GET /api/v1/categories/ category getCategories
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

	// swagger:route GET /api/v1/regions/ region getRegions
	//
	// Get all regions
	//
	//     Responses:
	//       200:     GetAllRegionsResponse
	//       default: GenericResponse
	v1.GET("/regions", c.RegionController.ReadAll)

	// swagger:route GET /api/v1/regions/:id region getOneRegion
	//
	// Get region by id
	//
	//     Responses:
	//       200:     GetOneRegionsResponse
	//       default: GenericResponse
	v1.GET("/regions/:id", c.RegionController.ReadOne)

	// swagger:route GET /api/v1/posts/:id posts getOnePost
	//
	// Get post by id
	//
	//     Responses:
	//       200:     GetOnePostResponse
	//       default: GenericResponse
	v1.GET("/posts/:id", c.PostController.ReadOne)

	// swagger:route GET /api/v1/posts/ posts getPosts
	//
	// Get all posts
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
	v1.POST("/posts", c.AuthMiddleware.IsLoggedIn, c.PostController.Create)

}
