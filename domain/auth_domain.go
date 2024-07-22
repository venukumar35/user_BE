package domain

import "TheBoys/app/model/response"

type AuthServices interface {
	SendOtp(email string) error
	VerifyOtp(otp string, email string) (*response.LoginResponse, error)
}
