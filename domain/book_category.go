package domain

import "gorm.io/gorm"

type BookCategory struct {
	gorm.Model
	BookID     uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
}
