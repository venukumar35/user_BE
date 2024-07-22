package routes

import (
	"TheBoys/api/middleware"
	"TheBoys/app/handler"
	"TheBoys/app/service"
	"TheBoys/domain"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoute(router *gin.RouterGroup, repo domain.ProductRepository, middleware *middleware.Middleware) {
	service := service.NewProductService(repo)
	handler := handler.NewProducthandler(service)

	routes := router.Group("/product")
	{
		routes.GET("/get", handler.GetProducts)
		routes.GET("/category", handler.GetCategorys)
		routes.GET("/byId", handler.GetProductById)
	}

}
