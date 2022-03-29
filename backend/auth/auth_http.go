package auth

import (
	"backend/roralis/email"
	"backend/roralis/jwt"
	"backend/roralis/otc"
	"backend/roralis/user"
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
