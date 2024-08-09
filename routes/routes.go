package routes

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/port/http"
	"bibliosphere_gin/service"
	"bibliosphere_gin/validators"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {

	registerUserRoutes(router, db, authMiddleware)
	registerBookRoutes(router, db, authMiddleware)
	registerBookCategoryRoutes(router, db, authMiddleware)
	registerLoanRoutes(router, db, authMiddleware)
	registerLibraryRoutes(router, db, authMiddleware)

	router.POST("/login_check", authMiddleware.LoginHandler)
}

func registerUserRoutes(router *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	userRepo := repositories.NewGormUserRepository(db)
	userValidator := validators.NewUserValidator()
	userService := service.NewUserService(userRepo, userValidator)
	userController := http.NewUserController(userService)

	userController.RegisterRoutes(router)
}

func registerBookRoutes(router *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	bookRepo := repositories.NewGormBookRepository(db)
	bookValidator := validators.NewBookValidator()
	bookService := service.NewBookService(bookRepo, bookValidator)
	bookController := http.NewBookController(bookService)

	bookController.RegisterRoutes(router)
}

func registerBookCategoryRoutes(router *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	bookCategoryRepo := repositories.NewGormBookCategoryRepository(db)
	bookCategoryValidator := validators.NewBookCategoryValidator()
	bookCategoryService := service.NewBookCategoryService(bookCategoryRepo, bookCategoryValidator)
	bookCategoryController := http.NewBookCategoryController(bookCategoryService)

	bookCategoryController.RegisterRoutes(router)
}

func registerLoanRoutes(router *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	loanRepo := repositories.NewGormLoanRepository(db)
	loanValidator := validators.NewLoanValidator()
	loanService := service.NewLoanService(loanRepo, loanValidator)
	loanController := http.NewLoanController(loanService)

	loanController.RegisterRoutes(router)
}

func registerLibraryRoutes(router *gin.Engine, db *gorm.DB, authMiddleware *jwt.GinJWTMiddleware) {
	libraryRepo := repositories.NewGormLibraryRepository(db)
	libraryValidator := validators.NewLibraryValidator()
	libraryService := service.NewLibraryService(libraryRepo, libraryValidator)
	libraryController := http.NewLibraryController(libraryService)

	libraryController.RegisterRoutes(router)
}
