package main

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/port/http"
	"bibliosphere_gin/service"
	"bibliosphere_gin/validators"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {

	userRepo := repositories.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := http.NewUserController(userService)

	bookRepo := repositories.NewGormBookRepository(db)
	bookValidator := validators.NewBookValidator()
	bookService := service.NewBookService(bookRepo, bookValidator)
	bookController := http.NewBookController(bookService)

	router := gin.Default()
	bookController.RegisterRoutes(router)
	userController.RegisterRoutes(router)

	return router
}
