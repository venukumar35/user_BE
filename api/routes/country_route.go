package routes

import (
	"TheBoys/app/handler"
	"TheBoys/app/service"
	"TheBoys/domain"

	"github.com/gin-gonic/gin"
)

func RegisterCountryRoute(router *gin.RouterGroup, repo domain.CountryRepository) {

	services := service.NewCountryService(repo)
	handler := handler.NewCountryHandler(services)
	routes := router.Group("/country")
	{
		routes.GET("", handler.FindCountry)
		routes.GET("/state", handler.FindState)
	}
}
