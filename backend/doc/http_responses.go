// This packages houses all swagger and other documtation stuff
package doc

import (
	"country/domain/entity"
)

// swagger:response GenericResponse
type GenericResponse struct {
	// Generic response, used when erros occurep
	// in: body
	Body entity.Response
}

// swagger:response GetAllUsersResponse
type GetAllUsersResponse struct {
	// Used when accesing GET /users
	// in: body
	Body []entity.User
}

// swagger:response GetOneUserResponse
type GetOneUserResponse struct {
	// Used when accesing GET /users
	// in: body
	Body entity.User
}

// swagger:response GetAllCategoriesResponse
type GetAllCategoriesResponse struct {
	// Used when accesing GET /users
	// in: body
	Body []entity.Category
}

// swagger:response GetOneCategoriesResponse
type GetOneCategoriesResponse struct {
	// Used when accesing GET /users
	// in: body
	Body entity.Category
}

// swagger:response GetAllRegionsResponse
type GetAllRegionsResponse struct {
	// Used when accesing GET /users
	// in: body
	Body []entity.Category
}

// swagger:response GetOneRegionsResponse
type GetOneRegionsResponse struct {
	// Used when accesing GET /users
	// in: body
	Body entity.Category
}

// swagger:response SignInSucces
type SignInSucces struct {
	// Used when accesing GET /users
	// in: body
	Body struct {
		Token string
	}
}

// swagger:response SignUpSucces
type SignUpSucces struct {
	// Used when accesing GET /users
	// in: body
	Body struct {
		User  entity.User
		Token string
	}
}

// swagger:response AboutMeSucces
type AboutMeSucces struct {
	// in: body
	Body struct {
		User entity.JWTClaims
	}
}
