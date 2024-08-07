package validators

import (
	"bibliosphere_gin/domain"

	"github.com/go-playground/validator/v10"
)


type LoanValidator interface {
	Validate(loan *domain.Loan) error
}

type loanValidator struct {
	validator *validator.Validate
}

func NewLoanValidator() LoanValidator {
	return &loanValidator{
		validator: validator.New(),
	}
}

func (v *loanValidator) Validate(loan *domain.Loan) error {
	return v.validator.Struct(loan)
}