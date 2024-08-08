package repositories

import (
	"bibliosphere_gin/domain"

	"gorm.io/gorm"
)

type LoanRepository interface {
	FindByID(id uint) (*domain.Loan, error)
	FindByUserID(id uint) (*domain.Loan, error)
	FindAll(loans *[]domain.Loan) error
	Save(loan *domain.Loan) error
	Delete(id uint) error
}

type GormLoanRepository struct {
	db *gorm.DB
}

func NewGormLoanRepository(db *gorm.DB) *GormLoanRepository {
	return &GormLoanRepository{db: db}
}

func (repo *GormLoanRepository) FindByID(id uint) (*domain.Loan, error) {
	var loan domain.Loan
	result := repo.db.First(&loan, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &loan, nil
}

func (repo *GormLoanRepository) FindByUserID(id uint) (*domain.Loan, error) {
	var loan domain.Loan
	result := repo.db.Where("user = ?", id).First(&loan)
	if result.Error != nil {
		return nil, result.Error
	}
	return &loan, nil
}

func (repo *GormLoanRepository) FindAll(loans *[]domain.Loan) error {
	result := repo.db.Find(loans)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormLoanRepository) Save(loan *domain.Loan) error {
	result := repo.db.Save(loan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *GormLoanRepository) Delete(id uint) error {
	result := repo.db.Delete(&domain.Loan{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
