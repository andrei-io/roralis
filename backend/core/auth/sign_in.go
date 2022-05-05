package auth

import (
	"backend/roralis/core/jwt"
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Request body for Sign In route
type signInRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Gin controller for sign-in flow
func (r *AuthController) SignIn(c *gin.Context) {
	// TODO
	var json signInRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, rest.Response{Message: err.Error()})
		return
	}

	user, err := r.userRepo.GetByEmail(json.Email)
	if err != nil {
		if errors.Is(err, repo.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, rest.NotFoundError)
			return
		} else {
			c.JSON(http.StatusInternalServerError, rest.Response{Message: err.Error()})
			return
		}

	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, rest.Response{Message: "Your password or email are incorrect"})
		return
	}

	payload := jwt.JWTClaims{
		ID:       user.ID,
		Name:     user.Name,
		Verified: user.Verified,
		Role:     user.Role,
	}
	token, err := r.jwtService.NewJWT(&payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, rest.Response{Message: "Your password or email are incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":    user.ID,
		"Token": token,
	})
}
