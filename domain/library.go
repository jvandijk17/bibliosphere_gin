package domain

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	Name       string `gorm:"size:255;not null"`
	Address    string `gorm:"size:255;not null"`
	City       string `gorm:"size:255;not null"`
	Province   string `gorm:"size:255;not null"`
	PostalCode string `gorm:"size:10;not null"`
	Users      []User `gorm:"foreignKey:LibraryID"`
}
