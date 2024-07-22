package handler

import (
	"TheBoys/api/response"
	"TheBoys/app/model/request"
	"TheBoys/domain"

	"github.com/gin-gonic/gin"
)

func NewProducthandler(service domain.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

type ProductHandler struct {
	services domain.ProductService
}

func (h ProductHandler) GetProducts(c *gin.Context) {
	var req request.RequestForGetProduct

	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	data, err := h.services.GetProducts()

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Product fetch successfully", data)
}
func (h ProductHandler) GetCategorys(c *gin.Context) {

	data, err := h.services.GetProductCategory()

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Product category fetch successfully", data)
}

func (h ProductHandler) GetProductById(c *gin.Context) {
	var req request.RequestProductById

	if err := c.ShouldBindQuery((&req)); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}
	data, err := h.services.GetProductById(req)

	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, "Product fetch successfully", data)

}
