package auth

import (
	"backend/roralis/core/jwt"
	"backend/roralis/core/user"
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Gin controller for signup
func (r *AuthController) SignUp(c *gin.Context) {

	var json user.User
	// Validate request form
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, rest.Response{Message: err.Error()})
		return
	}

	json.Role = 5
	json.Verified = false
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)

	if err != nil {
		// Failing to hash a password is a fatal error
		c.JSON(http.StatusInternalServerError, rest.Response{Message: err.Error()})
	}
	json.Password = string(hashedPassword)

	// Create in db. Will error out when invalid
	err = r.userRepo.Create(&json)
	if err != nil {
		if errors.Is(err, repo.ErrEmailTaken) {
			c.JSON(http.StatusConflict, rest.EmailTakenReponse)
			return
		} else {
			c.JSON(http.StatusInternalServerError, rest.Response{Message: err.Error()})
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
		c.JSON(http.StatusInternalServerError, rest.Response{Message: "Your password or email are incorrect"})
	}

	c.JSON(http.StatusOK, gin.H{
		"User":  json,
		"Token": token,
	})

}
