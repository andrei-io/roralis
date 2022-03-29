package category

import (
	"backend/roralis/dic"
	"backend/roralis/domain/entity"
	"backend/roralis/domain/repo/category"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Gin controller for GET /users
func ReadAll(c *gin.Context) {
	repo := dic.Container.Get(dic.CategoryRepo).(category.CategoryRepo)

	categories, err := repo.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)

}

// Gin controller for GET /users/:id
func ReadOne(c *gin.Context) {
	repo := dic.Container.Get(dic.CategoryRepo).(category.CategoryRepo)
	id := c.Param("id")

	category, err := repo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}
