package post

import (
	"country/dic"
	"country/domain/entity"
	"country/domain/repo/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin controller for creating a post, needs auth key
func Create(c *gin.Context) {
	postRepo := dic.Container.Get(dic.PostRepo).(post.IPostRepo)

	claimsRaw, exists := c.Get(dic.TokenKey)

	if !exists {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: "JWT claims object is missing"})
		return
	}

	claims, ok := claimsRaw.(*entity.JWTClaims)

	if !ok {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: "JWT claims is not of correct shape"})
		return
	}

	if claims.Role < 5 {
		c.JSON(http.StatusForbidden, entity.Response{Message: "Your email is not verified or you don't have enough permissions"})
		return
	}

	var json entity.Post

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Message: err.Error()})
		return
	}

	json.UserID = claims.ID

	err := postRepo.Create(&json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, json)

}
