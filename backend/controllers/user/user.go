package user

import (
	"backend/roralis/domain/repo/email"
	"backend/roralis/domain/repo/otc"
	"backend/roralis/domain/repo/user"
	"backend/roralis/domain/services/jwt"
)

type UserController struct {
	userRepo   user.UserRepo
	emailRepo  email.EmailRepo
	otcRepo    otc.OTCRepo
	jwtService jwt.JWTService
}

func NewUserController(u user.UserRepo, e email.EmailRepo, o otc.OTCRepo, j jwt.JWTService) UserController {
	return UserController{
		userRepo:   u,
		emailRepo:  e,
		otcRepo:    o,
		jwtService: j,
	}
}
