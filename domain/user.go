package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LibraryID        int       `gorm:"not null"`
	FirstName        string    `gorm:"size:255;not null"`
	LastName         string    `gorm:"size:255;not null"`
	Email            string    `gorm:"size:255;not null;unique"`
	Password         string    `gorm:"size:255;not null"`
	Address          string    `gorm:"size:255;not null"`
	City             string    `gorm:"size:255;not null"`
	Province         string    `gorm:"size:255;not null"`
	PostalCode       string    `gorm:"size:10;not null"`
	RegistrationDate time.Time `gorm:"not null"`
	BirthDate        *time.Time
	Reputation       int     `gorm:"not null"`
	Blocked          bool    `gorm:"not null"`
	Roles            string  `gorm:"type:longtext;not null"`
	Library          Library `gorm:"foreignKey:LibraryID"`
	// Assuming a relationship with the Loan table
	// Loans []Loan `gorm:"foreignKey:UserID"`
}
