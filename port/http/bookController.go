package http

import (
	"bibliosphere_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service service.BookService
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{service: service}
}

func (ctrl *BookController) RegisterRoutes(router *gin.Engine) {
	book := router.Group("/book")
	{
		book.GET("/", ctrl.GetAllBooks)
		book.GET("/:id", ctrl.GetBookByID)
		book.POST("/", ctrl.CreateOrUpdateBook)
		book.PUT("/:id", ctrl.CreateOrUpdateBook)
		book.DELETE("/:id", ctrl.DeleteBook)
	}
}

func (ctrl *BookController) GetAllBooks(context *gin.Context) {
	books, err := ctrl.service.GetAllBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, books)
}

func (ctrl *BookController) GetBookByID(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	book, err := ctrl.service.GetBookByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	context.JSON(http.StatusOK, book)
}

func (ctrl *BookController) CreateOrUpdateBook(context *gin.Context) {
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

	book, err := ctrl.service.CreateOrUpdateBook(id, data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	status := http.StatusCreated
	if id != nil {
		status = http.StatusOK
	}
	context.JSON(status, book)
}

func (ctrl *BookController) DeleteBook(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	err := ctrl.service.DeleteBook(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
