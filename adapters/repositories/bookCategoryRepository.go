package repositories

import (
	"bibliosphere_gin/domain"

	"gorm.io/gorm"
)

type BookCategoryRepository interface {
	FindByID(id uint) (*domain.BookCategory, error)
	FindAll(bookCategories *[]domain.BookCategory) error
	Save(bookCategory *domain.BookCategory) error
	Delete(id uint) error
}

type GormBookCategoryRepository struct {
	db *gorm.DB
}

func NewGormBookCategoryRepository(db *gorm.DB) *GormBookCategoryRepository {
	return &GormBookCategoryRepository{db: db}
}

func (repo *GormBookCategoryRepository) FindByID(id uint) (*domain.BookCategory, error) {
	var bookCategory domain.BookCategory
	result := repo.db.First(&bookCategory, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bookCategory, nil
}

func (repo *GormBookCategoryRepository) FindAll(bookCategories *[]domain.BookCategory) error {
	result := repo.db.Find(bookCategories)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormBookCategoryRepository) Save(bookCategory *domain.BookCategory) error {
	result := repo.db.Save(bookCategory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormBookCategoryRepository) Delete(id uint) error {
	result := repo.db.Delete(&domain.BookCategory{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
