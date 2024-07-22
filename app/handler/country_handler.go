package handler

import (
	"TheBoys/api/response"
	"TheBoys/app/model/request"
	"TheBoys/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewCountryHandler(service domain.CountryService) *CountryHandler {
	return &CountryHandler{service}
}

type CountryHandler struct {
	services domain.CountryService
}

func (h *CountryHandler) FindCountry(c *gin.Context) {
	var req request.CommonRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	data, err := h.services.FindCountry(req)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, "Country fetch successfully", data)

}

func (h *CountryHandler) FindState(c *gin.Context) {
	var req request.StateRequestBaseOnCountry
	fmt.Println("data from state request", req.CountryId)
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	data, err := h.services.FindState(req)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, "Country fetch successfully", data)

}
