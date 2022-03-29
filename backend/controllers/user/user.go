package user

import (
	"backend/roralis/auth"
	"backend/roralis/domain/repo/email"
	"backend/roralis/domain/repo/user"
	"backend/roralis/otc"
)

type UserController struct {
	userRepo    user.UserRepo
	emailRepo   email.EmailRepo
	otcRepo     otc.OTCRepo
	jwtService  auth.JWTService
	tokenString string
}

func NewUserController(u user.UserRepo, e email.EmailRepo, o otc.OTCRepo, j auth.JWTService, t string) UserController {
	return UserController{
		userRepo:    u,
		emailRepo:   e,
		otcRepo:     o,
		jwtService:  j,
		tokenString: t,
	}
}
