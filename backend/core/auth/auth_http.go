package auth

import (
	"backend/roralis/core/jwt"
	"backend/roralis/core/user"
)

type AuthController struct {
	userRepo    user.UserRepo
	jwtService  jwt.JWTService
	tokenString string
}

func NewAuthController(u user.UserRepo, j jwt.JWTService, t string) AuthController {
	return AuthController{
		userRepo:    u,
		jwtService:  j,
		tokenString: t,
	}
}
