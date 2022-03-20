package routes

import (
	"backend/roralis/dic"
	"backend/roralis/domain/entity"
	"backend/roralis/domain/services/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin middleware for checking if jwt is valid
// If valid will set a property on the context with the decoded claims
func IsLoggedIn(c *gin.Context) {
	jwtService := dic.Container.Get(dic.JWTService).(jwt.IJWTService)
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, entity.Response{Message: "Missing JWT token"})
		c.Abort()
		return
	}

	claims, err := jwtService.VerifyJWT(&token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.Response{Message: err.Error()})
		c.Abort()
		return
	}

	c.Set(dic.TokenKey, claims)

	c.Next()
}
