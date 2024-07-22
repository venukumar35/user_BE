package routes

import (
	"TheBoys/app/handler"
	"TheBoys/app/service"
	"TheBoys/domain"

	"github.com/gin-gonic/gin"
)

func RegisterNewUserRoute(router *gin.RouterGroup, repo domain.UserRepository) {
	service := service.NewUserServices(repo)

	handler := handler.NewUserHandler(service)

	routes := router.Group("/user")
	{
		routes.POST("/", handler.CreateUser)

	}
}
