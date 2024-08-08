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
	userValidator := validators.NewUserValidator()
	userService := service.NewUserService(userRepo, userValidator)
	userController := http.NewUserController(userService)

	bookRepo := repositories.NewGormBookRepository(db)
	bookValidator := validators.NewBookValidator()
	bookService := service.NewBookService(bookRepo, bookValidator)
	bookController := http.NewBookController(bookService)

	loanRepo := repositories.NewGormLoanRepository(db)
	loanValidator := validators.NewLoanValidator()
	loanService := service.NewLoanService(loanRepo, loanValidator)
	loanController := http.NewLoanController(loanService)

	libraryRepo := repositories.NewGormLibraryRepository(db)
	libraryValidator := validators.NewLibraryValidator()
	libraryService := service.NewLibraryService(libraryRepo, libraryValidator)
	libraryController := http.NewLibraryController(libraryService)

	router := gin.Default()
	bookController.RegisterRoutes(router)
	userController.RegisterRoutes(router)
	loanController.RegisterRoutes(router)
	libraryController.RegisterRoutes(router)

	return router
}
