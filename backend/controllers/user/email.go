package user

import (
	"backend/roralis/dic"
	"backend/roralis/domain/entity"
	"backend/roralis/domain/repo/email"
	"backend/roralis/domain/repo/otc"
	"backend/roralis/domain/repo/user"
	"backend/roralis/domain/services/jwt"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// GIN controller for GET /users/resend/:id
func ResendValidationEmail(c *gin.Context) {
	userRepo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
	emailRepo := dic.Container.Get(dic.EmailRepo).(email.IEmailRepo)
	otcRepo := dic.Container.Get(dic.OTCRepo).(otc.IOTCRepo)

	id := c.Param("id")

	user, err := userRepo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	if user.Verified {
		c.JSON(http.StatusOK, entity.Response{Message: "Email already confirmed"})
		return
	}

	verficationCode, err := entity.GenerateVerificationCode(6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	err = otcRepo.Set(user.ID, verficationCode, 30)
	if err != nil {
		// TODO: handle this better
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	if viper.GetString("ENV") == "PROD" {
		_, err := emailRepo.Send(user.Email, "backend/roralis Roads verification email", verficationCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
			return
		}
	} else {
		fmt.Printf("Verification code for user %s is %s\n", user.Email, verficationCode)
	}

	c.JSON(http.StatusOK, entity.SuccesResponse)

}

// Request body for the validate email route
type ValidateEmailRequest struct {
	Code string
}

func ValidateEmail(c *gin.Context) {
	userRepo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
	otcRepo := dic.Container.Get(dic.OTCRepo).(otc.IOTCRepo)

	jwtService := dic.Container.Get(dic.JWTService).(jwt.IJWTService)

	id := c.Param("id")

	var json ValidateEmailRequest
	// Validate request form
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Message: err.Error()})
		return
	}

	user, err := userRepo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	if user.Verified {
		c.JSON(http.StatusOK, entity.Response{Message: "Email already confirmed"})
		return
	}

	correctCode, err := otcRepo.Get(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}
	if json.Code != correctCode {
		c.JSON(http.StatusUnauthorized, entity.Response{Message: "The code is not correct"})
		return
	}

	user.Verified = true
	user.Role = 5

	err = userRepo.Update(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
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
		c.JSON(http.StatusInternalServerError, entity.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})

}
