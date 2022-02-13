package jwt

import (
	"errors"
	"fmt"
	"time"

	"country/domain/entity"

	"github.com/golang-jwt/jwt/v4"
)

type IJWTService interface {
	NewJWT(c *entity.JWTClaims) (string, error)
	VerifyJWT(token *string) (*entity.JWTClaims, error)
}

type JWTService struct {
	secret *entity.JWTSecret
}

// Constructor function
func NewJWTService(secret *entity.JWTSecret) *JWTService {
	return &JWTService{secret: secret}
}

// Returns a new JWT from the given claims
func (j *JWTService) NewJWT(c *entity.JWTClaims) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	// TODO: fix linting
	// nolint
	token.Claims = &entity.JWTInfo{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
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

// Verifies the given JWT and returns its custom claims
func (j *JWTService) VerifyJWT(token *string) (*entity.JWTClaims, error) {
	t, err := jwt.ParseWithClaims(*token, &entity.JWTInfo{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret.VerifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*entity.JWTInfo)

	if !ok || !t.Valid {
		return nil, errors.New("JWT is not valid or is of wrong shape")
	}

	return claims.JWTClaims, nil
}
