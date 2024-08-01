package service

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"bibliosphere_gin/utils"
	"bibliosphere_gin/validators"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	CreateOrUpdateUser(id *uint, data map[string]interface{}) (*domain.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo      repositories.UserRepository
	validator validators.UserValidator
}

func NewUserService(repo repositories.UserRepository, validator validators.UserValidator) UserService {
	return &userService{
		repo:      repo,
		validator: validator,
	}
}

func (userService *userService) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := userService.repo.FindAll(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService *userService) GetUserByID(id uint) (*domain.User, error) {
	return userService.repo.FindByID(id)
}

func (userService *userService) GetUserByEmail(email string) (*domain.User, error) {
	return userService.repo.FindByEmail(email)
}

func (userService *userService) CreateOrUpdateUser(id *uint, data map[string]interface{}) (*domain.User, error) {
	var user domain.User
	var err error
	if id != nil {
		existingUser, err := userService.repo.FindByID(*id)
		if err != nil {
			return nil, err
		}
		user = *existingUser
	}

	err = utils.AssignDataToStruct(&user, data)
	if err != nil {
		return nil, err
	}

	err = userService.validator.Validate(&user)
	if err != nil {
		return nil, err
	}

	err = userService.repo.Save(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userService *userService) DeleteUser(id uint) error {
	return userService.repo.Delete(id)
}
