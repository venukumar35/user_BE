package service

import (
	"TheBoys/app/model/response"
	"TheBoys/domain"
	"TheBoys/infrastructure/config"
	"TheBoys/utills"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func NewAuthServices(userRepo domain.UserRepository) domain.AuthServices {
	return &AuthServices{userRepo}
}

type AuthServices struct {
	userRepo domain.UserRepository
}

func (s *AuthServices) SendOtp(email string) error {

	userDetails, err := s.userRepo.FindUserByEmail(email)

	if err != nil {
		return err
	}
	if userDetails == nil {
		return errors.New("User email is not found")
	}
	currentTime := time.Now().UTC()

	otpExpiryTime := userDetails.OtpCreatedAt.Add(1 * time.Minute)

	if otpExpiryTime.After(currentTime) {
		return errors.New("Last otp is still in active state")
	}

	otp := utills.GenerateOTP()

	if len(otp) > 0 {

		err := s.userRepo.CreateOtp(otp, int16(userDetails.Id), userDetails.Email, userDetails.Mobile)

		if err != nil {
			return err
		}
	}

	err = utills.SendMail(userDetails.Email, otp, int(330))

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServices) VerifyOtp(otp string, email string) (*response.LoginResponse, error) {

	validateOtp, err := s.userRepo.FindUserByEmailWithOtp(email, otp)

	if !validateOtp {
		return nil, errors.New("Please enter the valid otp")
	}

	if err != nil {
		return nil, err
	}

	userDetails, err := s.userRepo.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if userDetails == nil {
		return nil, errors.New("Email is not founded")
	}
	token, err := generateToken(int16(userDetails.Id), userDetails.Email)
	fmt.Println("token", token)
	if err != nil {
		return nil, err
	}

	if len(token) > 0 {
		err = s.userRepo.UpdateWebToken(token, userDetails.Id)
		if err != nil {
			return nil, err
		}

		err = s.userRepo.UpdateLoginOtpStatus(userDetails.Id, otp)

		if err != nil {
			return nil, err
		}

		return &response.LoginResponse{
			Id:       userDetails.Id,
			Name:     userDetails.Name,
			Email:    userDetails.Email,
			Mobile:   userDetails.Mobile,
			WebToken: userDetails.WebToken,
			RoleId:   userDetails.RoleId,
			RoleName: userDetails.RoleName,
			IsActive: userDetails.IsActive,
		}, nil
	}

	return nil, errors.New("Token is not generated")
}

func generateToken(userId int16, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(config.Config.TokenDuration).Unix(),
		"is_web":  true,
	})
	return token.SignedString([]byte(config.Config.JwtSecretKey))
}
