package region

import (
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegionController struct {
	repo RegionRepo
}

func NewRegionController(c RegionRepo) RegionController {
	return RegionController{repo: c}
}

// Gin controller for GET /users
func (r *RegionController) ReadAll(c *gin.Context) {

	regions, err := r.repo.GetAll()

	if err != nil {
		if errors.Is(err, repo.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, rest.NotFoundResponse)
			return
		} else {
			c.JSON(http.StatusUnprocessableEntity, rest.Response{Message: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, regions)

}

// Gin controller for GET /users/:id
func (r *RegionController) ReadOne(c *gin.Context) {
	id := c.Param("id")

	region, err := r.repo.Get(id)

	if err != nil {
		if errors.Is(err, repo.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, rest.NotFoundResponse)
			return
		} else {
			c.JSON(http.StatusUnprocessableEntity, rest.Response{Message: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, region)
}
