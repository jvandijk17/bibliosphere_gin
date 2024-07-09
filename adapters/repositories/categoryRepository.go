package repositories

import (
	"bibliosphere_gin/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindByID(id uint) (*domain.Library, error)
	Save(library *domain.Library) error
	Delete(id uint) error
}

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (repo *GormCategoryRepository) FindByID(id uint) (*domain.Category, error) {
	var Category domain.Category
	result := repo.db.First(&Category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Category, nil
}

func (repo *GormCategoryRepository) Save(Category *domain.Category) error {
	result := repo.db.Save(Category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormCategoryRepository) Delete(id uint) error {
	result := repo.db.Delete(&domain.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
