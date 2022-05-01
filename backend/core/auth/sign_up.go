package auth

import (
	"backend/roralis/core/jwt"
	"backend/roralis/core/user"
	httpresponse "backend/roralis/shared/http_response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
)

// Gin controller for signup
func (r *AuthController) SignUp(c *gin.Context) {

	var json user.User
	// Validate request form
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, httpresponse.Response{Message: err.Error()})
		return
	}

	json.Role = 5
	json.Verified = false
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)

	if err != nil {
		// Failing to hash a password is a fatal error
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
	}
	json.Password = string(hashedPassword)

	// Create in db. Will error out when invalid
	err = r.userRepo.Create(&json)
	if err != nil {
		err := err.(*pgconn.PgError)
		message := err.Message
		if strings.Contains(message, "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, httpresponse.NewDuplicateEntityErrorResponse(err.ConstraintName))
			return
		} else {
			c.JSON(http.StatusUnprocessableEntity, httpresponse.Response{Message: err.Error()})
			return
		}
	}

	json.Password = "Secret"

	payload := jwt.JWTClaims{
		ID:       json.ID,
		Name:     json.Name,
		Verified: json.Verified,
		Role:     json.Role,
	}

	token, err := r.jwtService.NewJWT(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: "Your password or email are incorrect"})
	}

	c.JSON(http.StatusOK, gin.H{
		"User":  json,
		"Token": token,
	})

}
