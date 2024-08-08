package validators

import (
	"bibliosphere_gin/domain"

	"github.com/go-playground/validator/v10"
)

type BookCategoryValidator interface {
	Validate(bookCategory *domain.BookCategory) error
}

type bookCategoryValidator struct {
	validator *validator.Validate
}

func NewBookCategoryValidator() BookCategoryValidator {
	return &bookCategoryValidator{
		validator: validator.New(),
	}
}

func (v *bookCategoryValidator) Validate(bookCategory *domain.BookCategory) error {
	return v.validator.Struct(bookCategory)
}
