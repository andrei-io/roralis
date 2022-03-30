package auth

import (
	"backend/roralis/core/email"
	"backend/roralis/core/jwt"
	"backend/roralis/core/otc"
	"backend/roralis/core/user"
)

type AuthController struct {
	userRepo    user.UserRepo
	emailRepo   email.EmailRepo
	otcRepo     otc.OTCRepo
	jwtService  jwt.JWTService
	tokenString string
}

func NewAuthController(u user.UserRepo, e email.EmailRepo, o otc.OTCRepo, j jwt.JWTService, t string) AuthController {
	return AuthController{
		userRepo:    u,
		emailRepo:   e,
		otcRepo:     o,
		jwtService:  j,
		tokenString: t,
	}
}
