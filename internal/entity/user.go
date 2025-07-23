package entity

import (
	domainerr "github.com/NorthDice/ReflectDiary/internal/domain/errors"
	"net/mail"
	"regexp"
)

const (
	// MinPasswordLength defines the minimum allowed password length
	MinPasswordLength = 6

	// MaxPasswordLength defines the maximum allowed password length
	MaxPasswordLength = 32

	// MinUsernameLength defines the minimum allowed username length
	MinUsernameLength = 4

	// MaxUsernameLength defines the maximum allowed username length
	MaxUsernameLength = 50
)

// User represents the application user entity.
type User struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// ValidateEmail checks if the user's email is valid.
// Returns nil if valid.

func (u *User) ValidateEmail() error {
	if u.Email == IsEmptyString {
		return domainerr.ErrEmailIsRequired
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return domainerr.ErrInvalidEmail
	}

	return nil
}

// ValidateUsername checks that the username:
// - contains only latin letters, digits, and characters ! _ -
// - has a valid length.

func (u *User) ValidateUsername() error {
	re := regexp.MustCompile(`^[a-zA-Z0-9!_-]+$`)

	if u.Username == IsEmptyString {
		return domainerr.ErrUsernameIsRequired
	}
	if len(u.Username) < MinUsernameLength {
		return domainerr.ErrInvalidUsernameMinConstraint
	}
	if len(u.Username) > MaxUsernameLength {
		return domainerr.ErrInvalidUsernameMaxConstraint
	}
	if !re.MatchString(u.Username) {
		return domainerr.ErrInvalidUsernameConstraint
	}

	return nil
}

// ValidatePassword checks that the password:
// - is not empty,
// - has a valid length,
// - contains only allowed characters: latin letters, digits, and ! _ @ .

func (u *User) ValidatePassword() error {
	re := regexp.MustCompile(`^[a-zA-Z0-9!_@.]+$`)

	if u.Password == IsEmptyString {
		return domainerr.ErrPasswordIsRequired
	}
	if len(u.Password) < MinPasswordLength {
		return domainerr.ErrInvalidPasswordMinConstraint
	}
	if len(u.Password) > MaxPasswordLength {
		return domainerr.ErrInvalidPasswordMaxConstraint
	}
	if !re.MatchString(u.Password) {
		return domainerr.ErrInvalidPasswordConstraint
	}

	return nil
}
