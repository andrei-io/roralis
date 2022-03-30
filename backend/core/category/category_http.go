package category

import (
	httpresponse "backend/roralis/shared/http_response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	repo CategoryRepo
}

func NewCategoryController(c CategoryRepo) CategoryController {
	return CategoryController{repo: c}
}

// Gin controller for GET /users
func (r *CategoryController) ReadAll(c *gin.Context) {

	categories, err := r.repo.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)

}

// Gin controller for GET /users/:id
func (r *CategoryController) ReadOne(c *gin.Context) {
	id := c.Param("id")

	category, err := r.repo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}
