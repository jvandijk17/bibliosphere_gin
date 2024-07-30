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

	bookRepo := repositories.NewGormBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookController := http.NewBookController(bookService)

	router := gin.Default()
	bookController.RegisterRoutes(router)
	userController.RegisterRoutes(router)

	// Register other controllers similarly

	return router
}
