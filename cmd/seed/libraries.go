package seed

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedLibraries(db *gorm.DB) {

	libraryRepo := repositories.NewGormLibraryRepository(db)
	gofakeit.Seed(0)

	for i := 0; i < 5; i++ {
		randomLib := domain.Library{
			Name:       gofakeit.Company(),
			Address:    gofakeit.Address().Address,
			City:       gofakeit.City(),
			Province:   gofakeit.State(),
			PostalCode: gofakeit.Zip(),
		}
		if err := libraryRepo.Save(&randomLib); err != nil {
			log.Fatalf("Failed to seed random library %d: %v", i+1, err)
		}
	}

}
