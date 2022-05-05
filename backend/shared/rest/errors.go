package rest

// Standardized http response
type Response struct {
	Message string
}

var (
	NotFoundResponse  = Response{Message: "Record not found"}
	SuccesResponse    = Response{Message: "Succes"}
	EmailTakenReponse = Response{Message: "Email is already taken"}
)
