package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LibraryID        uint       `gorm:"not null" json:"libraryId" validate:"required"`
	FirstName        string     `gorm:"size:255;not null" json:"firstName" validate:"required,max=255"`
	LastName         string     `gorm:"size:255;not null" json:"lastName" validate:"required,max=255"`
	Email            string     `gorm:"size:255;not null;unique" json:"email" validate:"required,email"`
	Password         string     `gorm:"size:255;not null" json:"password" validate:"required,min=8,max=20,containsany=!@#$%^&*(),containsany=0123456789"`
	Address          string     `gorm:"size:255;not null" json:"address" validate:"required,max=255"`
	City             string     `gorm:"size:255;not null" json:"city" validate:"required,max=255"`
	Province         string     `gorm:"size:255;not null" json:"province" validate:"required,max=255"`
	PostalCode       string     `gorm:"size:10;not null" json:"postalCode" validate:"required,postalcode"`
	RegistrationDate time.Time  `gorm:"not null" json:"registrationDate" validate:"required,datetime=2006-01-02,lte"`
	BirthDate        *time.Time `json:"birthDate" validate:"required,datetime=2006-01-02,age_min=18"`
	Reputation       int        `gorm:"not null" json:"reputation" validate:"required,numeric,min=0"`
	Blocked          bool       `gorm:"not null" json:"blocked" validate:"required"`
	Roles            string     `gorm:"type:longtext;not null"`
	Library          Library    `gorm:"foreignKey:LibraryID"`
	// Assuming a relationship with the Loan table
	// Loans []Loan `gorm:"foreignKey:UserID"`
}
