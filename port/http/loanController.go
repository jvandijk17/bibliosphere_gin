package http

import (
	"bibliosphere_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	service service.LoanService
}

func NewLoanController(service service.LoanService) *LoanController {
	return &LoanController{service: service}
}

func (ctrl *LoanController) RegisterRoutes(router *gin.Engine) {
	loan := router.Group("/loan")
	{
		loan.GET("/", ctrl.GetAllLoans)
		loan.GET("/:id", ctrl.GetLoanByID)
		loan.GET("/:userId", ctrl.GetLoanByUserID)
		loan.POST("/", ctrl.CreateOrUpdateLoan)
		loan.PUT("/:id", ctrl.CreateOrUpdateLoan)
		loan.DELETE("/:id", ctrl.DeleteLoan)
	}
}

func (ctrl *LoanController) GetAllLoans(context *gin.Context) {
	loans, err := ctrl.service.GetAllLoans()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, loans)
}

func (ctrl *LoanController) GetLoanByID(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	loan, err := ctrl.service.GetLoanByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}
	context.JSON(http.StatusOK, loan)
}

func (ctrl *LoanController) GetLoanByUserID(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("userId"), 10, 32)
	loan, err := ctrl.service.GetLoanByUserID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}
	context.JSON(http.StatusOK, loan)
}

func (ctrl *LoanController) CreateOrUpdateLoan(context *gin.Context) {
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

	loan, err := ctrl.service.CreateOrUpdateLoan(id, data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	status := http.StatusCreated
	if id != nil {
		status = http.StatusOK
	}
	context.JSON(status, loan)
}

func (ctrl *LoanController) DeleteLoan(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	err := ctrl.service.DeleteLoan(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
