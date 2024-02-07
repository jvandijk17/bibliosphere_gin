package service

import (
	adapters "bibliosphere_gin/adapters/repositories"
	"bibliosphere_gin/domain"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo adapters.UserRepository
}

func NewAuthService(repo adapters.UserRepository) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) AuthenticateUser(email, password string, jwtKey []byte) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := generateJWTToken(user, jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func checkPasswordHash(password, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateJWTToken(user *domain.User, jwtKey []byte) (string, error) {

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil

}
