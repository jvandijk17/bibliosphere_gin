package repositories

import (
	"bibliosphere_gin/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindByID(id uint) (*domain.Book, error)
	FindAll(books *[]domain.Book) error
	Save(book *domain.Book) error
	Delete(id uint) error
}

type GormBookRepository struct {
	db *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{db: db}
}

func (repo *GormBookRepository) FindByID(id uint) (*domain.Book, error) {
	var book domain.Book
	result := repo.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (repo *GormBookRepository) FindAll(books *[]domain.Book) error {
	result := repo.db.Find(books)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormBookRepository) Save(book *domain.Book) error {
	result := repo.db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormBookRepository) Delete(id uint) error {
	result := repo.db.Delete(&domain.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
