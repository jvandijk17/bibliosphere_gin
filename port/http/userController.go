package http

import (
	"bibliosphere_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) RegisterRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/", ctrl.GetAllUsers)
		user.GET("/:id", ctrl.GetUserByID)
		user.GET("/:email", ctrl.GetUserByEmail)
		user.POST("/", ctrl.CreateOrUpdateUser)
		user.PUT("/:id", ctrl.CreateOrUpdateUser)
		user.DELETE("/:id", ctrl.DeleteUser)
	}
}

func (ctrl *UserController) GetAllUsers(context *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUserByID(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	user, err := ctrl.service.GetUserByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUserByEmail(context *gin.Context) {
	user, err := ctrl.service.GetUserByEmail(context.Param("email"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (ctrl *UserController) CreateOrUpdateUser(context *gin.Context) {
	idParam := context.Param("id")
	var id *uint
	if idParam != "" {
		idVal, _ := strconv.ParseUint(idParam, 10, 32)
		id = new(uint)
		*id = uint(idVal)
	}

	var data map[string]interface{}
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := ctrl.service.CreateOrUpdateUser(id, data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	status := http.StatusCreated
	if id != nil {
		status = http.StatusOK
	}
	context.JSON(status, user)
}

func (ctrl *UserController) DeleteUser(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	err := ctrl.service.DeleteUser(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
