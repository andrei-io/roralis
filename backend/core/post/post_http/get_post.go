package posthttp

import (
	"backend/roralis/core/post"
	httpresponse "backend/roralis/shared/http_response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	repo     post.PostRepo
	tokenKey string
}

func NewPostController(r post.PostRepo, t string) PostController {
	return PostController{repo: r, tokenKey: t}
}

// Gin controller for reading all posts
func (r *PostController) ReadAll(c *gin.Context) {

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "200"))

	user_id := c.Query("user_id")
	if user_id != "" {
		posts, err := r.repo.GetByUserID(user_id)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
		return
	}

	posts, err := r.repo.GetAll(offset, limit, true)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// Gin controller for reading a single post
func (r *PostController) ReadOne(c *gin.Context) {

	id := c.Param("id")

	post, err := r.repo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
