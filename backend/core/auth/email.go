package auth

import (
	"backend/roralis/core/jwt"
	"backend/roralis/core/otc"
	httpresponse "backend/roralis/shared/http_response"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// GIN controller for GET /users/resend/:id
func (r *AuthController) ResendValidationEmail(c *gin.Context) {

	id := c.Param("id")

	user, err := r.userRepo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
		return
	}

	if user.Verified {
		c.JSON(http.StatusOK, httpresponse.Response{Message: "Email already confirmed"})
		return
	}

	verficationCode, err := otc.GenerateVerificationCode(6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
		return
	}

	err = r.otcRepo.Set(user.ID, verficationCode, 30)
	if err != nil {
		// TODO: handle this better
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
		return
	}

	if viper.GetString("ENV") == "PROD" {
		_, err := r.emailRepo.Send(user.Email, "backend/roralis Roads verification email", verficationCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
			return
		}
	} else {
		fmt.Printf("Verification code for user %s is %s\n", user.Email, verficationCode)
	}

	c.JSON(http.StatusOK, httpresponse.SuccesResponse)

}

// Request body for the validate email route
type validateEmailRequest struct {
	Code string
}

func (r *AuthController) ValidateEmail(c *gin.Context) {

	id := c.Param("id")

	var json validateEmailRequest
	// Validate request form
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, httpresponse.Response{Message: err.Error()})
		return
	}

	user, err := r.userRepo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, httpresponse.NotFoundError)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
		return
	}

	if user.Verified {
		c.JSON(http.StatusOK, httpresponse.Response{Message: "Email already confirmed"})
		return
	}

	correctCode, err := r.otcRepo.Get(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
		return
	}
	if json.Code != correctCode {
		c.JSON(http.StatusUnauthorized, httpresponse.Response{Message: "The code is not correct"})
		return
	}

	user.Verified = true
	user.Role = 5

	err = r.userRepo.Update(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
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
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})

}
