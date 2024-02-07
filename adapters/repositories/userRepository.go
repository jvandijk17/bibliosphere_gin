package repositories

import (
	"bibliosphere_gin/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Save(user *domain.User) error
	Delete(id uint) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (repo *GormUserRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *GormUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *GormUserRepository) Save(user *domain.User) error {
	result := repo.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormUserRepository) Delete(id uint) error {
	result := repo.db.Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
