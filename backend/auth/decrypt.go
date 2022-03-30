package auth

import (
	"backend/roralis/jwt"
	httpresponse "backend/roralis/shared/http_response"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin controller that decodes and sends jwt back
func (r *AuthController) AboutMe(c *gin.Context) {

	claimsRaw, exists := c.Get(r.tokenString)

	if !exists {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: "JWT claims object is missing"})
		return
	}

	claims, ok := claimsRaw.(*jwt.JWTClaims)

	if !ok {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: "JWT claims is not of correct shape"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": claims,
	})

}
