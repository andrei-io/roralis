package user

import (
	"backend/roralis/domain/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	u.Password = "Secret"
	u.Email = "Secret"

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}
