package service

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"bibliosphere_gin/utils"
	"bibliosphere_gin/validators"
)

type LibraryService interface {
	GetAllLibraries() ([]domain.Library, error)
	GetLibraryById(id uint) (*domain.Library, error)
	CreateOrUpdateLibrary(id *uint, data map[string]interface{}) (*domain.Library, error)
	DeleteLibrary(id uint) error
}

type libraryService struct {
	repo      repositories.LibraryRepository
	validator validators.LibraryValidator
}

func NewLibraryService(repo repositories.LibraryRepository, validator validators.LibraryValidator) LibraryService {
	return &libraryService{
		repo:      repo,
		validator: validator,
	}
}

func (libraryService *libraryService) GetAllLibraries() ([]domain.Library, error) {
	var libraries []domain.Library
	err := libraryService.repo.FindAll(&libraries)
	if err != nil {
		return nil, err
	}
	return libraries, nil
}

func (libraryService *libraryService) GetLibraryById(id uint) (*domain.Library, error) {
	return libraryService.repo.FindByID(id)
}

func (libraryService *libraryService) CreateOrUpdateLibrary(id *uint, data map[string]interface{}) (*domain.Library, error) {
	var library domain.Library
	var err error
	if id != nil {
		existingLibrary, err := libraryService.repo.FindByID(*id)
		if err != nil {
			return nil, err
		}
		library = *existingLibrary
	}

	err = utils.AssignDataToStruct(&library, data)
	if err != nil {
		return nil, err
	}

	err = libraryService.validator.Validate(&library)
	if err != nil {
		return nil, err
	}

	err = libraryService.repo.Save(&library)
	if err != nil {
		return nil, err
	}
	return &library, nil
}

func (libraryService *libraryService) DeleteLibrary(id uint) error {
	return libraryService.repo.Delete(id)
}
