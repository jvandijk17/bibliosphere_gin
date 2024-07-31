package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	LibraryID       uint       `gorm:"not null" json:"libraryId" validate:"required"`
	Title           string     `gorm:"size:255;not null" json:"title" validate:"required"`
	Author          string     `gorm:"size:255;not null" json:"author" validate:"required"`
	Publisher       string     `gorm:"size:255;not null" json:"publisher" validate:"required"`
	ISBN            string     `gorm:"size:13;not null" json:"isbn" validate:"required,isbn"`
	PublicationYear string     `gorm:"type:date;not null" json:"publication_year" validate:"required,datetime=2006-01-02"`
	PageCount       uint       `gorm:"not null" json:"pageCount" validate:"required,gt=0"`
	Categories      []Category `gorm:"many2many:book_categories;"`
	Loans           []Loan     `gorm:"foreignKey:BookID"`
}
