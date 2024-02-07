package http

import (
	"bibliosphere_gin/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) RegisterRoutes(router *gin.Engine) {
	router.GET("/users", uc.GetUsers)
	router.POST("/users", uc.CreateUser)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	// Implementación del método
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// Implementación del método
}
