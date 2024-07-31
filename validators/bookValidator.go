package validators

import (
	"bibliosphere_gin/domain"

	"github.com/go-playground/validator/v10"
)

type BookValidator interface {
	Validate(book *domain.Book) error
}

type bookValidator struct {
	validator *validator.Validate
}

func NewBookValidator() BookValidator {
	return &bookValidator{
		validator: validator.New(),
	}
}

func (v *bookValidator) Validate(book *domain.Book) error {
	return v.validator.Struct(book)
}
