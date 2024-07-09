package seed

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedBooks(db *gorm.DB) {
	bookRepo := repositories.NewGormBookRepository(db)
	libraries := getLibraries(db)

	for i := 0; i < 5; i++ {
		randomBook := createRandomBook(libraries)
		if err := bookRepo.Save(&randomBook); err != nil {
			log.Fatalf("Failed to seed random book %d: %v", i+1, err)
		}
	}
}

func getLibraries(db *gorm.DB) []domain.Library {
	var libraries []domain.Library
	if err := db.Find(&libraries).Error; err != nil {
		log.Fatalf("Failed to retrieve libraries: %v", err)
	}

	if len(libraries) == 0 {
		log.Fatalf("No libraries found. Seed the libraries first.")
	}
	return libraries
}

func createRandomBook(libraries []domain.Library) domain.Book {
	gofakeit.Seed(0)
	randomLibrary := libraries[rand.Intn(len(libraries))]
	return domain.Book{
		LibraryID:       randomLibrary.ID,
		Title:           gofakeit.BookTitle(),
		Author:          gofakeit.BookAuthor(),
		Publisher:       gofakeit.Company(),
		ISBN:            gofakeit.Generate("#############"),
		PublicationYear: gofakeit.Date().Format("2006-01-02"),
		PageCount:       gofakeit.Number(100, 1000),
	}
}
