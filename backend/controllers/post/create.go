package post

import (
	"backend/roralis/dic"
	"backend/roralis/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin controller for creating a post, needs auth key
func (r *PostController) Create(c *gin.Context) {

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

	err := r.repo.Create(&json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, json)

}
