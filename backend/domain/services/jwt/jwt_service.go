// Contains everything related to working with JWTs: service
package jwt

import (
	"errors"
	"fmt"
	"time"

	"backend/roralis/domain/entity"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	NewJWT(c *entity.JWTClaims) (string, error)
	VerifyJWT(token *string) (*entity.JWTClaims, error)
}

type jwtService struct {
	secret *entity.JWTSecret
}

// Check interface at compile time
var _ JWTService = (*jwtService)(nil)

// Constructor function
func NewJWTService(secret *entity.JWTSecret) *jwtService {
	return &jwtService{secret: secret}
}

// Returns a new JWT from the given claims
func (j *jwtService) NewJWT(c *entity.JWTClaims) (string, error) {
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
func (j *jwtService) VerifyJWT(token *string) (*entity.JWTClaims, error) {
	t, err := jwt.ParseWithClaims(*token, &entity.JWTInfo{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// the library doc is not very clear about what i should be returning but this works
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
