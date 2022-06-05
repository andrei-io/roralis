package user

import (
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo UserRepo
}

func NewUserController(u UserRepo) UserController {
	return UserController{userRepo: u}
}

// Gin controller for GET /users/:id
func (r *UserController) ReadOne(c *gin.Context) {
	id := c.Param("id")

	u, err := r.userRepo.Get(id)
	if errors.Is(err, repo.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, rest.NotFoundResponse)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, rest.Response{Message: err.Error()})
		return
	}

	u.Password = "Secret"
	u.Email = "Secret"

	c.JSON(http.StatusOK, u)
}
