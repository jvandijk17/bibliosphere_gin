package http

import (
	"bibliosphere_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LibraryController struct {
	service service.LibraryService
}

func NewLibraryController(service service.LibraryService) *LibraryController {
	return &LibraryController{service: service}
}

func (ctrl *LibraryController) RegisterRoutes(router *gin.Engine) {
	library := router.Group("/library")
	{
		library.GET("/", ctrl.GetAllLibraries)
		library.GET("/:id", ctrl.GetLibraryById)
		library.POST("/", ctrl.CreateOrUpdateLibrary)
		library.PUT("/:id", ctrl.CreateOrUpdateLibrary)
		library.DELETE("/:id", ctrl.DeleteLibrary)
	}
}

func (ctrl *LibraryController) GetAllLibraries(context *gin.Context) {
	libraries, err := ctrl.service.GetAllLibraries()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, libraries)
}

func (ctrl *LibraryController) GetLibraryById(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	library, err := ctrl.service.GetLibraryById(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
		return
	}
	context.JSON(http.StatusOK, library)
}

func (ctrl *LibraryController) CreateOrUpdateLibrary(context *gin.Context) {
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

	library, err := ctrl.service.CreateOrUpdateLibrary(id, data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	status := http.StatusCreated
	if id != nil {
		status = http.StatusOK
	}
	context.JSON(status, library)
}

func (ctrl *LibraryController) DeleteLibrary(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	err := ctrl.service.DeleteLibrary(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
