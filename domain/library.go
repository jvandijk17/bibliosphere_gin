package domain

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	Name       string `gorm:"size:255;not null" json:"name" validate:"required"`
	Address    string `gorm:"size:255;not null" json:"address" validate:"required"`
	City       string `gorm:"size:255;not null" json:"city" validate:"required"`
	Province   string `gorm:"size:255;not null" json:"province" validate:"required"`
	PostalCode string `gorm:"size:10;not null" json:"postalCode" validate:"required,gt=0"`
	Users      []User `gorm:"foreignKey:LibraryID"`
}
