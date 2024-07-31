package service

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"bibliosphere_gin/validators"
	"time"
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

	err = s.assignDataToBook(&book, data)
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

func (s *bookService) assignDataToBook(book *domain.Book, data map[string]interface{}) error {
	if title, ok := data["title"].(string); ok {
		book.Title = title
	}
	if author, ok := data["author"].(string); ok {
		book.Author = author
	}
	if publisher, ok := data["publisher"].(string); ok {
		book.Publisher = publisher
	}
	if isbn, ok := data["isbn"].(string); ok {
		book.ISBN = isbn
	}
	if publicationYear, ok := data["publication_year"].(string); ok {
		parsedDate, err := time.Parse("2006-01-02", publicationYear)
		if err != nil {
			return err
		}
		book.PublicationYear = parsedDate.Format("2006-01-02")
	}
	if pageCount, ok := data["page_count"].(uint); ok {
		book.PageCount = pageCount
	}
	if libraryID, ok := data["libraryId"].(uint); ok {
		book.LibraryID = libraryID
	}

	return nil
}
