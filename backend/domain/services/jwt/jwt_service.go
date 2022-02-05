package jwt

import (
	"errors"
	"time"

	"country/domain/entity"

	"github.com/golang-jwt/jwt/v4"
)

type IJWTService interface {
	NewJWT(c *entity.JWTClaims) (string, error)
	VerifyJWT(token *string) error
}

type JWTService struct {
	secret *entity.JWTSecret
}

func NewJWTService(secret *entity.JWTSecret) *JWTService {
	return &JWTService{secret: secret}
}

func (j *JWTService) NewJWT(c *entity.JWTClaims) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	// TODO: fix linting
	// nolint
	token.Claims = &entity.JWTInfo{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		&entity.JWTClaims{
			Name:     c.Name,
			ID:       c.ID,
			Verified: c.Verified,
			Role:     c.Role,
		},
	}

	return token.SignedString(j.secret.SignKey)
}

func (j *JWTService) VerifyJWT(token *string) error {
	return errors.New("Not implemented yet")
}
