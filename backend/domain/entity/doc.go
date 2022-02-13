// This packages contains all entities.
// Entities are things/structs that exist outside the infrastructure.
// Basically business logic.
package entity

func init() {
	NotFoundError = Response{Message: "Record not found"}
	SuccesResponse = Response{Message: "Succes"}
}
