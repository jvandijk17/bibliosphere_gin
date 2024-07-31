package service

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"bibliosphere_gin/utils"
	"bibliosphere_gin/validators"
)

type BookService interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id uint) (*domain.Book, error)
	CreateOrUpdateBook(id *uint, data map[string]interface{}) (*domain.Book, error)
	DeleteBook(id uint) error
}

type bookService struct {
	repo      repositories.BookRepository
	validator validators.BookValidator
}

func NewBookService(repo repositories.BookRepository, validator validators.BookValidator) BookService {
	return &bookService{
		repo:      repo,
		validator: validator,
	}
}

func (s *bookService) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book
	err := s.repo.FindAll(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *bookService) GetBookByID(id uint) (*domain.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateOrUpdateBook(id *uint, data map[string]interface{}) (*domain.Book, error) {
	var book domain.Book
	var err error
	if id != nil {
		existingBook, err := s.repo.FindByID(*id)
		if err != nil {
			return nil, err
		}
		book = *existingBook
	}

	err = utils.AssignDataToStruct(&book, data)
	if err != nil {
		return nil, err
	}

	err = s.validator.Validate(&book)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
