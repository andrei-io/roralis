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

	// swagger:route GET /users/ user getUser
	//
	// Get all users
	//
	//     Responses:
	//       default: GenericResponse
	//       404:
	//       200:     GetAllUsersResponse
	app.GET("/users", user.ReadAll)

	// swagger:route GET /users/:id user getUser
	//
	// Get one user
	//
	//     Responses:
	//       default: GenericResponse
	//       200:     GetOneUserResponse
	app.GET("/users/:id", user.ReadOne)

	// swagger:route PUT /users/:id user updateUser
	//
	// Update user
	//
	//     Responses:
	//       default: GenericResponse
	app.PUT("/users/:id", user.Update)

	// swagger:route POST /users/ user createUser
	//
	// Create user
	//
	//     Responses:
	//       default: GenericResponse
	app.POST("/users", user.Create)

	// swagger:route DELETE /users/:id user deleteUser
	//
	// Delete user
	//
	//     Responses:
	//       default: GenericResponse
	app.DELETE("/users/:id", user.Delete)

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
