package service

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"bibliosphere_gin/utils"
	"bibliosphere_gin/validators"
)

type LoanService interface {
	GetAllLoans() ([]domain.Loan, error)
	GetLoanByID(id uint)(*domain.Loan, error)
	GetLoanByUserID(id uint)(*domain.Loan, error)
	CreateOrUpdateLoan(id *uint, data map[string]interface{}) (*domain.Loan, error)
	DeleteLoan(id uint) error
}

type loanService struct {
	repo repositories.LoanRepository
	validator validators.LoanValidator
}

func NewLoanService(repo repositories.LoanRepository, validator validators.LoanValidator) LoanService {
	return &loanService{
		repo: repo,
		validator: validator,
	}
}

func (loanService *loanService) GetAllLoans() ([]domain.Loan, error) {
	var loans []domain.Loan
	err := loanService.repo.FindAll(&loans)
	if err != nil {
		 return nil, err
	}
	return loans, err
}

func (loanService *loanService) GetLoanByID(id uint) (*domain.Loan, error) {
	return loanService.repo.FindByID(id)
}

func (loanService *loanService) GetLoanByUserID(id uint) (*domain.Loan, error) {
	return loanService.repo.FindByUserID(id)
}

func (loanService *loanService) CreateOrUpdateLoan(id *uint, data map[string]interface{}) (*domain.Loan, error) {
	var loan domain.Loan
	var err error
	if id != nil {
		existingLoan, err := loanService.repo.FindByID(*id)
		if err != nil {
			return nil, err
		}
		loan = *existingLoan
	}

	err = utils.AssignDataToStruct(&loan, data)
	if err != nil {
		return nil, err
	}
	
	err = loanService.validator.Validate(&loan)
	if err != nil {
		return nil, err
	}

	err = loanService.repo.Save(&loan)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (loanService *loanService) DeleteLoan(id uint) error {
	return loanService.repo.Delete(id)
}