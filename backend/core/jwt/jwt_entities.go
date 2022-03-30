package jwt

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Name     string `json:"Name"`
	ID       uint64 `json:"ID"`
	Verified bool   `json:"Verified"`
	Role     uint8  `json:"Role"`
}

type JWTInfo struct {
	*jwt.StandardClaims
	*JWTClaims
}

type JWTSecret struct {
	SignKey   *rsa.PrivateKey
	VerifyKey *rsa.PublicKey
}
