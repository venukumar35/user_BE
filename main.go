package main

import (
	schedulers "TheBoys/Schedulers"
	"TheBoys/api/routes"
	"TheBoys/infrastructure/config"
	"TheBoys/infrastructure/database"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	loadConfiguration()
	setupServer()
}

func loadConfiguration() {
	err := config.Load()

	if err != nil {
		panic(err)
	}
}

func setupCors(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	router.Use(cors.New(corsConfig))
}

func setupRateLimiter(router *gin.Engine) {
	limiter := tollbooth.NewLimiter(30, nil)
	router.Use(tollbooth_gin.LimitHandler(limiter))
}

func setupServer() {
	db, pDb, err := database.IntDB()

	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	setupRateLimiter(router)
	setupCors(router)
	setupRoutes(router, db, pDb)

	Schedule := schedulers.NewSchedule(db, pDb)

	Schedule.ProductSchedulers()

	router.Static("/public", "./public")

	err = router.Run(config.Config.Port)

	if err != nil {
		panic(err)
	}
}
func setupRoutes(router *gin.Engine, db *gorm.DB, pDb *gorm.DB) {
	routes.SetupRoutes(router, db)

}
