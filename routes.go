package main

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/middleware"
	"bibliosphere_gin/routes"
	"bibliosphere_gin/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	setupCORS(router)
	userRepo := repositories.NewGormUserRepository(db)
	authService := service.NewAuthService(userRepo)

	authMiddleware, err := middleware.JWTMiddleware(authService)
	if err != nil {
		panic("JWT Middleware initialization failed: " + err.Error())
	}

	routes.RegisterRoutes(router, db, authMiddleware)

	return router
}

func setupCORS(router *gin.Engine) {
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))
}
