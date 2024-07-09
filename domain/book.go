package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	LibraryID       uint       `gorm:"not null"`
	Title           string     `gorm:"size:255;not null"`
	Author          string     `gorm:"size:255;not null"`
	Publisher       string     `gorm:"size:255;not null"`
	ISBN            string     `gorm:"size:13;not null"`
	PublicationYear string     `gorm:"type:date;not null"`
	PageCount       int        `gorm:"not null"`
	Categories      []Category `gorm:"many2many:book_categories;"`
	Loans           []Loan     `gorm:"foreignKey:BookID"`
}
