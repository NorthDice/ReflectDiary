package user

import (
	"fmt"
	"github.com/NorthDice/ReflectDiary/internal/entity"
	"github.com/NorthDice/ReflectDiary/internal/usecase/interfaces"
)

// RegisterRequest represents the data required to register a new user.
type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterResponse represents the data returned after a successful user registration.
type RegisterResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// RegisterUseCase handles the business logic for user registration.
type RegisterUseCase struct {
	userRepository  interfaces.UserRepository
	passwordService interfaces.PasswordService
	authService     interfaces.AuthService
}

// NewRegisterUseCase creates a new instance of RegisterUseCase.
func NewRegisterUseCase(
	userRepository interfaces.UserRepository,
	passwordService interfaces.PasswordService,
	authService interfaces.AuthService,
) *RegisterUseCase {
	return &RegisterUseCase{
		userRepository:  userRepository,
		passwordService: passwordService,
		authService:     authService,
	}
}

// Register handles the full registration flow:
// 1. Validates user input,
// 2. Checks if the email already exists,
// 3. Hashes the password,
// 4. Saves the new user to the repository,
// 5. Generates an auth token,
// 6. Returns the resulting user info with token.
func (uc *RegisterUseCase) Register(req RegisterRequest) (*RegisterResponse, error) {
	user := &entity.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}

	if err := user.ValidateUsername(); err != nil {
		return nil, err
	}

	if err := user.ValidatePassword(); err != nil {
		return nil, err
	}

	existingUser, err := uc.userRepository.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	hashedPassword, err := uc.passwordService.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	savedUserId, err := uc.userRepository.Save(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := uc.authService.GenerateToken(savedUserId)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	response := &RegisterResponse{
		ID:       savedUserId,
		Email:    req.Email,
		Username: req.Username,
		Token:    token,
	}

	return response, nil
}
