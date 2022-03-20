package post

import (
	"backend/roralis/dic"
	"backend/roralis/domain/entity"
	"backend/roralis/domain/repo/post"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Gin controller for reading a single post
func ReadOne(c *gin.Context) {
	postRepo := dic.Container.Get(dic.PostRepo).(post.IPostRepo)

	id := c.Param("id")

	post, err := postRepo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
