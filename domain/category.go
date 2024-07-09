package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Books []Book `gorm:"many2many:book_categories;"`
}
