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
	"country/controllers/region"
	"country/controllers/user"
	_ "country/doc" // for swagger responses

	"github.com/gin-gonic/gin"
)

// Mount the routes
func MountRoutes(app *gin.Engine) {
	// swagger:route GET /users/:id user getUser
	//
	// Get one user
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GetOneUserResponse
	app.GET("/users/:id", user.ReadOne)

	// swagger:route POST /users/signup user signup
	//
	// Sign Up
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GenericResponse
	app.POST("/users/signup", user.SignUp)

	// swagger:route GET /categories/ category getCategory
	//
	// Get all categories
	//
	//     Responses:
	//       200:     GetAllCategoriesResponse
	//       default: GenericResponse
	app.GET("/categories", category.ReadAll)

	// swagger:route GET /categories/:id category getCategory
	//
	// Get category by id
	//
	//     Responses:
	//       200:     GetOneCategoriesResponse
	//       default: GenericResponse
	app.GET("/categories/:id", category.ReadOne)

	// swagger:route GET /regions/ region getRegion
	//
	// Get all regions
	//
	//     Responses:
	//       200:     GetAllRegionsResponse
	//       default: GenericResponse
	app.GET("/regions", region.ReadAll)

	// swagger:route GET /regions/:id region getCategory
	//
	// Get region by id
	//
	//     Responses:
	//       200:     GetOneRegionsResponse
	//       default: GenericResponse
	app.GET("/regions/:id", region.ReadOne)

}
