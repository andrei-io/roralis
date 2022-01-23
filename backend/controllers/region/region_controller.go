package region

import (
	"country/dic"
	"country/domain/entity"
	"country/domain/repo/region"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Gin controller for GET /users
func ReadAll(c *gin.Context) {
	repo := dic.Container.Get(dic.RegionRepo).(region.IRegionRepo)

	regions, err := repo.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, regions)

}

// Gin controller for GET /users/:id
func ReadOne(c *gin.Context) {
	repo := dic.Container.Get(dic.RegionRepo).(region.IRegionRepo)
	id := c.Param("id")

	region, err := repo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, region)
}
