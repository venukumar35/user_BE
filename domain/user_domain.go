package domain

import (
	"TheBoys/app/model/request"
	"TheBoys/app/model/response"
)

type UserService interface {
	CreateUser(req request.UserCreationRequest) error
}

type UserRepository interface {
	FindUserByEmail(email string) (*response.LoginResponse, error)
	CreateOtp(otp string, userId int16, email string, mobile string) error
	FindUserByEmailWithOtp(email string, otp string) (bool, error)
	UpdateWebToken(token string, userId uint) error
	UpdateLoginOtpStatus(userId uint, otp string) error
	FindUserById(userId uint) (*response.LoginResponse, error)
	CreateUser(req request.UserCreationRequest) error
}
