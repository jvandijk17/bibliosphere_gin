package seed

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedBookCategories(db *gorm.DB) {
	bookCategoryRepo := repositories.NewGormBookCategoryRepository(db)
	books := getBooks(db)
	categories := getCategories(db)

	for i := 0; i < 5; i++ {
		randomBookCategory := createRandomBookCategory(books, categories)
		if err := bookCategoryRepo.Save(&randomBookCategory); err != nil {
			log.Fatalf("Failed to seed random book category %d: %v", i+1, err)
		}
	}

}

func getBooks(db *gorm.DB) []domain.Book {
	var books []domain.Book
	if err := db.Find(&books).Error; err != nil {
		log.Fatalf("Failed to retrieve books: %v", err)
	}
	if len(books) == 0 {
		log.Fatalf("No books found. Seed the books first.")
	}
	return books
}

func getCategories(db *gorm.DB) []domain.Category {
	var categories []domain.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Fatalf("Failed to retrieve categories: %v", err)
	}
	if len(categories) == 0 {
		log.Fatalf("No categories found. Seed the categories first.")
	}
	return categories
}

func createRandomBookCategory(books []domain.Book, categories []domain.Category) domain.BookCategory {
	gofakeit.Seed(0)
	randomBook := books[rand.Intn(len(books))]
	randomCategory := categories[rand.Intn(len(categories))]

	return domain.BookCategory{
		BookID:     randomBook.ID,
		CategoryID: randomCategory.ID,
	}
}
