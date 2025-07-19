package interfaces

import (
	"github.com/NorthDice/ReflectDiary/internal/entity"
	"github.com/NorthDice/ReflectDiary/internal/usecase/user"
)

// UserPresenter defines methods for formatting user-related responses,
// such as registration and login results.
type UserPresenter interface {
	// RegisterResponse formats the response data after a successful user registration.
	// Takes the registered User entity and a token string.
	// Returns a RegisterResponse struct used for presentation.
	RegisterResponse(user *entity.User, token string) *user.RegisterResponse

	// LoginResponse formats the response data after a successful user login.
	// Takes the authenticated User entity and a token string.
	// Returns a LoginResponse struct used for presentation.
	LoginResponse(user *entity.User, token string) *user.LoginResponse
}
