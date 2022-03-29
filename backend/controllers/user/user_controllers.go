package user

import (
	"backend/roralis/dic"
	"backend/roralis/domain/entity"
	"backend/roralis/domain/repo/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Gin controller for GET /users/:id
func ReadOne(c *gin.Context) {
	repo := dic.Container.Get(dic.UserRepo).(user.UserRepo)
	id := c.Param("id")

	u, err := repo.Get(id)
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
