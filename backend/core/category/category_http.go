package category

import (
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if err != nil {
		if errors.Is(err, repo.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, rest.NotFoundResponse)
			return
		} else {
			c.JSON(http.StatusUnprocessableEntity, rest.Response{Message: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, categories)

}

// Gin controller for GET /users/:id
func (r *CategoryController) ReadOne(c *gin.Context) {
	id := c.Param("id")

	category, err := r.repo.Get(id)

	if err != nil {
		if errors.Is(err, repo.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, rest.NotFoundResponse)
			return
		} else {
			c.JSON(http.StatusUnprocessableEntity, rest.Response{Message: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, category)
}
