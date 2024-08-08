package http

import (
	"bibliosphere_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookCategoryController struct {
	service service.BookCategoryService
}

func NewBookCategoryController(service service.BookCategoryService) *BookCategoryController {
	return &BookCategoryController{service: service}
}

func (ctrl *BookCategoryController) RegisterRoutes(router *gin.Engine) {
	bookCategory := router.Group("/bookCategory")
	{
		bookCategory.GET("/", ctrl.GetAllBookCategories)
		bookCategory.GET("/:id", ctrl.GetBookCategoryByID)
		bookCategory.POST("/", ctrl.CreateOrUpdateBookCategory)
		bookCategory.PUT("/:id", ctrl.CreateOrUpdateBookCategory)
		bookCategory.DELETE("/:id", ctrl.DeleteBookCategory)
	}
}

func (ctrl *BookCategoryController) GetAllBookCategories(context *gin.Context) {
	bookCategories, err := ctrl.service.GetAllBookCategories()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, bookCategories)
}

func (ctrl *BookCategoryController) GetBookCategoryByID(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	bookCategory, err := ctrl.service.GetBookCategoryByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "BookCategory not found"})
		return
	}
	context.JSON(http.StatusOK, bookCategory)
}

func (ctrl *BookCategoryController) CreateOrUpdateBookCategory(context *gin.Context) {
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

	bookCategory, err := ctrl.service.CreateOrUpdateBookCategory(id, data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	status := http.StatusCreated
	if id != nil {
		status = http.StatusOK
	}
	context.JSON(status, bookCategory)
}

func (ctrl *BookCategoryController) DeleteBookCategory(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	err := ctrl.service.DeleteBookCategory(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "BookCategory not found"})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
