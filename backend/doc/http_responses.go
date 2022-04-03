// This packages houses all swagger and other documtation stuff
package doc

import (
	"backend/roralis/core/category"
	"backend/roralis/core/jwt"
	"backend/roralis/core/post"
	"backend/roralis/core/region"
	"backend/roralis/core/user"
	httpresponse "backend/roralis/shared/http_response"
)

// swagger:response GenericResponse
type GenericResponse struct {
	// Generic response, used when erros occurep
	// in: body
	Body httpresponse.Response
}

// swagger:response GetOneUserResponse
type GetOneUserResponse struct {
	// Used when accesing GET /users
	// in: body
	Body user.User
}

// swagger:response GetAllCategoriesResponse
type GetAllCategoriesResponse struct {
	// Used when accesing GET /users
	// in: body
	Body []category.Category
}

// swagger:response GetOneCategoriesResponse
type GetOneCategoriesResponse struct {
	// Used when accesing GET /users
	// in: body
	Body category.Category
}

// swagger:response GetAllRegionsResponse
type GetAllRegionsResponse struct {
	// Used when accesing GET /users
	// in: body
	Body []region.Region
}

// swagger:response GetOneRegionsResponse
type GetOneRegionsResponse struct {
	// Used when accesing GET /users
	// in: body
	Body region.Region
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
		User  user.User
		Token string
	}
}

// swagger:response AboutMeSucces
type AboutMeSucces struct {
	// in: body
	Body struct {
		User jwt.JWTClaims
	}
}

// swagger:response GetOnePostResponse
type GetOnePostResponse struct {
	// in: body
	Body post.Post
}

// swagger:response GetAllPostResponse
type GetAllPostResponse struct {
	// in: body
	Body []post.Post
}
