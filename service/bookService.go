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

func (bookService *bookService) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book
	err := bookService.repo.FindAll(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (bookService *bookService) GetBookByID(id uint) (*domain.Book, error) {
	return bookService.repo.FindByID(id)
}

func (bookService *bookService) CreateOrUpdateBook(id *uint, data map[string]interface{}) (*domain.Book, error) {
	var book domain.Book
	var err error
	if id != nil {
		existingBook, err := bookService.repo.FindByID(*id)
		if err != nil {
			return nil, err
		}
		book = *existingBook
	}

	err = utils.AssignDataToStruct(&book, data)
	if err != nil {
		return nil, err
	}

	err = bookService.validator.Validate(&book)
	if err != nil {
		return nil, err
	}

	err = bookService.repo.Save(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (bookService *bookService) DeleteBook(id uint) error {
	return bookService.repo.Delete(id)
}
