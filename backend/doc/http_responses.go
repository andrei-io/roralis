package doc

import "country/domain/entity"

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
