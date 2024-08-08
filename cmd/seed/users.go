package seed

import (
	"bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/config"
	"bibliosphere_gin/domain"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	userRepo := repositories.NewGormUserRepository(db)

	var library domain.Library
	if err := db.First(&library).Error; err != nil {
		log.Fatalf("Failed to fetch a library for seeding users: %v", err)
	}

	hashedPasswordStr := hashPassword(config.AppConfig.JwtTestPass)

	jwtUser := createUser(library.ID, "Test", "User", config.AppConfig.JwtTestMail, hashedPasswordStr, "ROLE_USER,ROLE_ADMIN", nil)
	if err := userRepo.Save(&jwtUser); err != nil {
		log.Fatalf("Failed to seed JWT user: %v", err)
	}

	for i := 0; i < 19; i++ {
		email := strings.ToLower(faker.Email())
		randomUser := createUser(library.ID, faker.FirstName(), faker.LastName(), email, hashedPasswordStr, "ROLE_USER", randomBirthDate())
		if err := userRepo.Save(&randomUser); err != nil {
			log.Fatalf("Failed to seed random user %d: %v", i+1, err)
		}
	}
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hashedPassword)
}

func createUser(libraryID uint, firstName, lastName, email, password, roles string, birthDate *time.Time) domain.User {
	return domain.User{
		LibraryID:  libraryID,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Password:   password,
		Address:    gofakeit.Address().Address,
		City:       gofakeit.City(),
		Province:   gofakeit.State(),
		PostalCode: gofakeit.Zip(),
		BirthDate:  birthDate,
		Reputation: 0,
		Roles:      roles,
		Blocked:    false,
	}
}

func randomBirthDate() *time.Time {
	now := time.Now()
	maxAge := 60
	minAge := 20

	years := rand.Intn(maxAge-minAge+1) + minAge
	days := rand.Intn(365)

	randomDate := now.AddDate(-years, 0, -days)
	return &randomDate
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
