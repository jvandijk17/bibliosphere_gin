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

func (ctrl *BookController) GetAllBooks(c *gin.Context) {
	books, err := ctrl.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (ctrl *BookController) GetBookByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := ctrl.service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) CreateOrUpdateBook(c *gin.Context) {
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

	book, err := ctrl.service.CreateOrUpdateBook(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	status := http.StatusCreated
	if id != nil {
		status = http.StatusOK
	}
	c.JSON(status, book)
}

func (ctrl *BookController) DeleteBook(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := ctrl.service.DeleteBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
