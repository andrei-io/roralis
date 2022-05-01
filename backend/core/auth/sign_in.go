package auth

import (
	"backend/roralis/core/jwt"
	httpresponse "backend/roralis/shared/http_response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
		c.JSON(http.StatusBadRequest, httpresponse.Response{Message: err.Error()})
		return
	}

	user, err := r.userRepo.GetByEmail(json.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, httpresponse.Response{Message: "Your password or email are incorrect"})
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
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: "Your password or email are incorrect"})
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}
