package seed

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	categoryRepo := repositories.NewGormCategoryRepository(db)

	for i := 0; i < 5; i++ {
		randomCategory := createRandomCategory()
		if err := categoryRepo.Save(&randomCategory); err != nil {
			log.Fatalf("Failed to seed random category %d: %v", i+1, err)
		}
	}

}

func createRandomCategory() domain.Category {
	gofakeit.Seed(0)
	return domain.Category{
		Name: gofakeit.BookGenre(),
	}
}
