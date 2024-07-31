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
		user.POST("/", ctrl.CreateOrUpdateUser)
		user.PUT("/:id", ctrl.CreateOrUpdateUser)
		user.DELETE("/:id", ctrl.DeleteUser)
	}
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := ctrl.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) CreateOrUpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	var id *uint
	if idParam != "" {
		idVal, _ := strconv.ParseUint(idParam, 10, 32)
		id = new(uint)
		*id = uint(idVal)
	}

	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := ctrl.service.CreateOrUpdateUser(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	status := http.StatusCreated
	if id != nil {
		status = http.StatusOK
	}
	c.JSON(status, user)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := ctrl.service.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
