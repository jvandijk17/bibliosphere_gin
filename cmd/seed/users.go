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

	jwtUser := domain.User{
		Email:    os.Getenv("JWT_TEST_MAIL"),
		Password: os.Getenv("JWT_TEST_PASS"),
		Roles:    "ROLE_USER,ROLE_ADMIN",
		Blocked:  false,
	}
	if err := userRepo.Save(&jwtUser); err != nil {
		log.Fatalf("Failed to seed JWT user: %v", err)
	}

	for i := 0; i < 19; i++ {
		randomUser := domain.User{
			Email:    faker.Email(),
			Password: os.Getenv("JWT_TEST_PASS"),
			Roles:    "ROLE_USER",
			Blocked:  false,
		}
		if err := userRepo.Save(&randomUser); err != nil {
			log.Fatalf("Failed to seed random user %d: %v", i+1, err)
		}
	}

}
