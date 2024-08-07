package domain

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"userId" validate:"required"`
	BookID     uint   `gorm:"not null" json:"bookId" validate:"required"`
	LoanDate   string `gorm:"type:date;not null" json:"loanDate" validate:"required, datetime=2006-01-02"`
	ReturnDate string `gorm:"type:date"`
}
