package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	NewJWT(c *JWTClaims) (string, error)
	VerifyJWT(token *string) (*JWTClaims, error)
}

type jwtService struct {
	secret *JWTSecret
}

// Check interface at compile time
var _ JWTService = (*jwtService)(nil)

// Constructor function
func NewJWTService(secret *JWTSecret) *jwtService {
	return &jwtService{secret: secret}
}

// Returns a new JWT from the given claims
func (j *jwtService) NewJWT(c *JWTClaims) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	// TODO: fix linting
	// nolint
	token.Claims = &JWTInfo{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
		&JWTClaims{
			Name:     c.Name,
			ID:       c.ID,
			Verified: c.Verified,
			Role:     c.Role,
		},
	}

	return token.SignedString(j.secret.SignKey)
}

// Verifies the given JWT and returns its custom claims
func (j *jwtService) VerifyJWT(token *string) (*JWTClaims, error) {
	t, err := jwt.ParseWithClaims(*token, &JWTInfo{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// the library doc is not very clear about what i should be returning but this works
		return j.secret.VerifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*JWTInfo)

	if !ok || !t.Valid {
		return nil, errors.New("JWT is not valid or is of wrong shape")
	}

	return claims.JWTClaims, nil
}
