package seed

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"log"
	"os"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {

	userRepo := repositories.NewGormUserRepository(db)

	var library domain.Library
	if err := db.First(&library).Error; err != nil {
		log.Fatalf("Failed to fetch a library for seeding users: %v", err)
	}

	jwtUser := domain.User{
		LibraryID: int(library.ID),
		Email:     os.Getenv("JWT_TEST_MAIL"),
		Password:  os.Getenv("JWT_TEST_PASS"),
		Roles:     "ROLE_USER,ROLE_ADMIN",
		Blocked:   false,
	}
	if err := userRepo.Save(&jwtUser); err != nil {
		log.Fatalf("Failed to seed JWT user: %v", err)
	}

	for i := 0; i < 19; i++ {
		randomUser := domain.User{
			LibraryID: int(library.ID),
			Email:     faker.Email(),
			Password:  os.Getenv("JWT_TEST_PASS"),
			Roles:     "ROLE_USER",
			Blocked:   false,
		}
		if err := userRepo.Save(&randomUser); err != nil {
			log.Fatalf("Failed to seed random user %d: %v", i+1, err)
		}
	}

}
