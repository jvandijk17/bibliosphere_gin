package repositories

import (
	"bibliosphere_gin/domain"

	"gorm.io/gorm"
)

type BookCategoryRepository interface {
	FindByID(id uint) (*domain.Library, error)
	Save(library *domain.Library) error
	Delete(id uint) error
}

type GormBookCategoryRepository struct {
	db *gorm.DB
}

func NewGormBookCategoryRepository(db *gorm.DB) *GormBookCategoryRepository {
	return &GormBookCategoryRepository{db: db}
}

func (repo *GormBookCategoryRepository) FindByID(id uint) (*domain.BookCategory, error) {
	var BookCategory domain.BookCategory
	result := repo.db.First(&BookCategory, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &BookCategory, nil
}

func (repo *GormBookCategoryRepository) Save(BookCategory *domain.BookCategory) error {
	result := repo.db.Save(BookCategory)
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
