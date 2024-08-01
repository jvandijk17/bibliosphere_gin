package repositories

import (
	"bibliosphere_gin/domain"

	"gorm.io/gorm"
)

type LibraryRepository interface {
	FindByID(id uint) (*domain.Library, error)
	FindAll(libraries *[]domain.Library) error
	Save(library *domain.Library) error
	Delete(id uint) error
}

type GormLibraryRepository struct {
	db *gorm.DB
}

func NewGormLibraryRepository(db *gorm.DB) *GormLibraryRepository {
	return &GormLibraryRepository{db: db}
}

func (repo *GormLibraryRepository) FindByID(id uint) (*domain.Library, error) {
	var library domain.Library
	result := repo.db.First(&library, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &library, nil
}

func (repo *GormLibraryRepository) FindAll(libraries *[]domain.Library) error {
	result := repo.db.Find(libraries)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormLibraryRepository) Save(library *domain.Library) error {
	result := repo.db.Save(library)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormLibraryRepository) Delete(id uint) error {
	result := repo.db.Delete(&domain.Library{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
