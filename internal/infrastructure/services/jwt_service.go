package services

import (
	"context"
	"github.com/NorthDice/ReflectDiary/internal/infrastructure/repository/postgres"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthService struct {
	userRepository *postgres.UserRepository
}

func NewAuthService(userRepository *postgres.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (service *AuthService) GenerateToken(ctx context.Context, userID int) (string, error) {
	user, err := service.userRepository.FindById(ctx, userID)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) ValidateToken(ctx context.Context, token string) (string, error) {
	return "", nil
}
