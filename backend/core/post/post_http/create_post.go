package posthttp

import (
	"backend/roralis/core/jwt"
	"backend/roralis/core/post"
	"backend/roralis/shared/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin controller for creating a post, needs jwt key
func (r *PostController) Create(c *gin.Context) {

	claimsRaw, exists := c.Get(r.tokenKey)

	if !exists {
		c.JSON(http.StatusInternalServerError, rest.Response{Message: "JWT claims object is missing"})
		return
	}

	claims, ok := claimsRaw.(*jwt.JWTClaims)

	if !ok {
		c.JSON(http.StatusInternalServerError, rest.Response{Message: "JWT claims is not of correct shape"})
		return
	}

	if claims.Role < 5 {
		c.JSON(http.StatusForbidden, rest.Response{Message: "Your email is not verified or you don't have enough permissions"})
		return
	}

	var json post.Post

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, rest.Response{Message: err.Error()})
		return
	}

	json.UserID = claims.ID

	err := r.repo.Create(&json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, rest.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, json)

}
