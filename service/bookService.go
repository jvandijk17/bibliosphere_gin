package service

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
)

type BookService interface {
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id uint) (*domain.Book, error)
	CreateOrUpdateBook(id *uint, data map[string]interface{}) (*domain.Book, error)
	DeleteBook(id uint) error
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
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
	if id != nil {
		existingBook, err := s.repo.FindByID(*id)
		if err != nil {
			return nil, err
		}
		book = *existingBook
	}

	// Update book fields based on data
	// Assuming data contains valid book fields
	// For simplicity, data validation and mapping are omitted

	book.Title = data["title"].(string)
	book.Author = data["author"].(string)
	// ... (other fields)

	err := s.repo.Save(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
