package validators

import (
	"bibliosphere_gin/domain"

	"github.com/go-playground/validator/v10"
)

type UserValidator interface {
	Validate(user *domain.User) error
}

type userValidator struct {
	validator *validator.Validate
}

func NewUserValidator() UserValidator {
	return &userValidator{
		validator: validator.New(),
	}
}

func (v *userValidator) Validate(user *domain.User) error {
	return v.validator.Struct(user)
}
