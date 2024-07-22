package handler

import (
	"TheBoys/api/response"
	"TheBoys/app/model/request"
	"TheBoys/domain"

	"github.com/gin-gonic/gin"
)

func NewAuthHandler(services domain.AuthServices) *AuthHandler {
	return &AuthHandler{services}
}

type AuthHandler struct {
	services domain.AuthServices
}

func (h AuthHandler) SendOtp(c *gin.Context) {

	var req request.SendotpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	err := h.services.SendOtp(req.Email)

	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}
	response.Success(c, "Successfully  otp send to email", nil)
}
func (h AuthHandler) VerifyOtp(c *gin.Context) {

	var req request.VerifyAuthOtp

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	data, err := h.services.VerifyOtp(req.Otp, req.Email)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Login successfully completed", data)

}
