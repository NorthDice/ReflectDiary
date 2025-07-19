package user

import (
	"fmt"
	"github.com/NorthDice/ReflectDiary/internal/entity"
	"github.com/NorthDice/ReflectDiary/internal/usecase/interfaces"
	"strings"
)

// LoginRequest represents the data required to authenticate a user.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the data returned after a successful login.
type LoginResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

// LoginUseCase handles the business logic for user login.
type LoginUseCase struct {
	userRepository  interfaces.UserRepository
	passwordService interfaces.PasswordService
	authService     interfaces.AuthService
}

// NewLoginUseCase creates a new instance of LoginUseCase.
func (lu *LoginUseCase) NewLoginUseCase(
	userRepository interfaces.UserRepository,
	passwordService interfaces.PasswordService,
	authService interfaces.AuthService,
) *LoginUseCase {
	return &LoginUseCase{
		userRepository:  userRepository,
		passwordService: passwordService,
		authService:     authService,
	}
}

// Login authenticates a user based on provided credentials.
// It checks the email and password, compares the password hash,
// generates a token if valid, and returns a login response.
func (lu *LoginUseCase) Login(req LoginRequest) (*LoginResponse, error) {
	if strings.TrimSpace(req.Email) == entity.IsEmptyString {
		return nil, fmt.Errorf("email is required")
	}

	if strings.TrimSpace(req.Password) == entity.IsEmptyString {
		return nil, fmt.Errorf("password is required")
	}
	user, err := lu.userRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if err = lu.passwordService.ComparePassword(req.Password, user.Password); err != nil {
		return nil, fmt.Errorf("password wrong")
	}

	token, err := lu.authService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}

	return response, nil
}
