package entity

import "fmt"

type Response struct {
	Message string
}

var NotFoundError Response
var SuccesResponse Response

func init() {
	NotFoundError = Response{Message: "Record not found"}
	SuccesResponse = Response{Message: "Succes"}
}

func NewDuplicateEntityErrorResponse(field string) Response {
	return Response{Message: fmt.Sprintf("This value is already taken: %s", field)}
}
