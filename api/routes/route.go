package routes

import (
	"TheBoys/api/middleware"
	"TheBoys/app/service"
	"TheBoys/infrastructure/config"
	"TheBoys/infrastructure/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s server running in port no: %s", config.Config.Name, config.Config.Port)})
	})

	apiRouter := router.Group("/api")

	userRepo := repository.NewUserRepository(db)
	countryRepo := repository.NewCountryRepository(db)
	productRepo := repository.NewProductRepository(db)

	authServices := service.NewAuthServices(userRepo)
	middleWare := middleware.NewMiddleware(userRepo)

	fmt.Println("mid", middleWare)

	RegisterNewAuthRoute(apiRouter, authServices)
	RegisterNewUserRoute(apiRouter, userRepo)
	RegisterCountryRoute(apiRouter, countryRepo)
	RegisterProductRoute(apiRouter, productRepo, middleWare)

}
