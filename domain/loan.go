package domain

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	UserID     uint   `gorm:"not null"`
	BookID     uint   `gorm:"not null"`
	LoanDate   string `gorm:"type:date;not null"`
	ReturnDate string `gorm:"type:date"`
}
