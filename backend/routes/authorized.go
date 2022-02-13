package routes

import (
	"country/dic"
	"country/domain/entity"
	"country/domain/services/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin middleware for checking if jwt is valid
// If valid will set a property on the context with the decoded claims
func IsLoggedIn(c *gin.Context) {
	jwtService := dic.Container.Get(dic.JWTService).(jwt.IJWTService)
	token := c.GetHeader("Authorization")

	claims, err := jwtService.VerifyJWT(&token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.Response{Message: err.Error()})
		c.Abort()
		return
	}

	c.Set(dic.TokenKey, claims)

	c.Next()
}
