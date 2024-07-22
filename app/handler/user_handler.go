package handler

import (
	"TheBoys/api/response"
	"TheBoys/app/model/request"
	"TheBoys/domain"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service}
}

type UserHandler struct {
	service domain.UserService
}

func (h UserHandler) CreateUser(c *gin.Context) {
	var req request.UserCreationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	err := h.service.CreateUser(req)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Account created successfully", nil)
}
