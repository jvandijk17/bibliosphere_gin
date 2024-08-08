package service

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"bibliosphere_gin/utils"
	"bibliosphere_gin/validators"
)

type BookCategoryService interface {
	GetAllBookCategories() ([]domain.BookCategory, error)
	GetBookCategoryByID(id uint) (*domain.BookCategory, error)
	CreateOrUpdateBookCategory(id *uint, data map[string]interface{}) (*domain.BookCategory, error)
	DeleteBookCategory(id uint) error
}

type bookCategoryService struct {
	repo      repositories.BookCategoryRepository
	validator validators.BookCategoryValidator
}

func NewBookCategoryService(repo repositories.BookCategoryRepository, validator validators.BookCategoryValidator) BookCategoryService {
	return &bookCategoryService{
		repo:      repo,
		validator: validator,
	}
}

func (bookCategoryService *bookCategoryService) GetAllBookCategories() ([]domain.BookCategory, error) {
	var bookCategories []domain.BookCategory
	err := bookCategoryService.repo.FindAll(&bookCategories)
	if err != nil {
		return nil, err
	}
	return bookCategories, nil
}

func (bookCategoryService *bookCategoryService) GetBookCategoryByID(id uint) (*domain.BookCategory, error) {
	return bookCategoryService.repo.FindByID(id)
}

func (bookCategoryService *bookCategoryService) CreateOrUpdateBookCategory(id *uint, data map[string]interface{}) (*domain.BookCategory, error) {
	var bookCategory domain.BookCategory
	var err error
	if id != nil {
		existingBookCategory, err := bookCategoryService.repo.FindByID(*id)
		if err != nil {
			return nil, err
		}
		bookCategory = *existingBookCategory
	}

	err = utils.AssignDataToStruct(&bookCategory, data)
	if err != nil {
		return nil, err
	}

	err = bookCategoryService.validator.Validate(&bookCategory)
	if err != nil {
		return nil, err
	}

	err = bookCategoryService.repo.Save(&bookCategory)
	if err != nil {
		return nil, err
	}
	return &bookCategory, nil
}

func (bookCategoryService *bookCategoryService) DeleteBookCategory(id uint) error {
	return bookCategoryService.repo.Delete(id)
}
