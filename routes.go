package main

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/port/http"
	"bibliosphere_gin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	userRepo := repositories.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := http.NewUserController(userService)

	router := gin.Default()
	userController.RegisterRoutes(router)
	return router
}
