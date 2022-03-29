package user

import (
	"backend/roralis/dic"
	"backend/roralis/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin controller that decodes and sends jwt back
func (r *UserController) AboutMe(c *gin.Context) {

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

	c.JSON(http.StatusOK, gin.H{
		"User": claims,
	})

}
