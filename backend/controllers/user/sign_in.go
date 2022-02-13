package user

import (
	"country/dic"
	"country/domain/entity"
	"country/domain/repo/user"
	"country/domain/services/jwt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Request body for Sign In route
type SignInRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Gin controller for sign-in flou
func SignIn(c *gin.Context) {
	// TODO
	jwtService := dic.Container.Get(dic.JWTService).(jwt.IJWTService)
	userRepo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
	var json SignInRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Message: err.Error()})
		return
	}

	user, err := userRepo.GetByEmail(json.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.Response{Message: "Your password or email are incorrect"})
		return
	}

	payload := entity.JWTClaims{
		ID:       user.ID,
		Name:     user.Name,
		Verified: user.Verified,
		Role:     user.Role,
	}
	token, err := jwtService.NewJWT(&payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: "Your password or email are incorrect"})
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}
