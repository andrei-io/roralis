package middleware

import (
	"backend/roralis/jwt"
	httpresponse "backend/roralis/shared/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	jwtService  jwt.JWTService
	tokenString string
}

func NewAuthService(j jwt.JWTService, t string) AuthService {
	return AuthService{jwtService: j, tokenString: t}
}

// Gin middleware for checking if jwt is valid
// If valid will set a property on the context with the decoded claims
func (r *AuthService) IsLoggedIn(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, httpresponse.Response{Message: "Missing JWT token"})
		c.Abort()
		return
	}

	claims, err := r.jwtService.VerifyJWT(&token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, httpresponse.Response{Message: err.Error()})
		c.Abort()
		return
	}

	c.Set(r.tokenString, claims)

	c.Next()
}
