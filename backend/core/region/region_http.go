package region

import (
	httpresponse "backend/roralis/shared/http_response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, regions)

}

// Gin controller for GET /users/:id
func (r *RegionController) ReadOne(c *gin.Context) {
	id := c.Param("id")

	region, err := r.repo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, region)
}
