package validators

import (
	"bibliosphere_gin/domain"
	"github.com/go-playground/validator/v10"
)

type LibraryValidator interface {
	Validate(library *domain.Library) error
}

type libraryValidator struct {
	validator *validator.Validate
}

func NewLibraryValidator() LibraryValidator {
	return &libraryValidator{
		validator: validator.New(),
	}
}

func (v *libraryValidator) Validate(library *domain.Library) error {
	return v.validator.Struct(library)
}
