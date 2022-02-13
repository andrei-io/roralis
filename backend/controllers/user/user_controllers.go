package user

import (
	"country/dic"
	"country/domain/entity"
	"country/domain/repo/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Gin controller for GET /users/:id
func ReadOne(c *gin.Context) {
	repo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
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
