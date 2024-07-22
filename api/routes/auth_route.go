package routes

import (
	"TheBoys/app/handler"
	"TheBoys/domain"

	"github.com/gin-gonic/gin"
)

func RegisterNewAuthRoute(routes *gin.RouterGroup, service domain.AuthServices) {

	handler := handler.NewAuthHandler(service)
	authRoute := routes.Group("/auth")
	{
		authRoute.POST("", handler.SendOtp)
		authRoute.POST("/verify", handler.VerifyOtp)
	}
}
