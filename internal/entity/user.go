package entity

import (
	"fmt"
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
		return fmt.Errorf("email is required")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return fmt.Errorf("invalid email format")
	}

	return nil
}

// ValidateUsername checks that the username:
// - contains only latin letters, digits, and characters ! _ -
// - has a valid length.

func (u *User) ValidateUsername() error {
	re := regexp.MustCompile(`^[a-zA-Z0-9!_-]+$`)

	if u.Username == IsEmptyString {
		return fmt.Errorf("username is required")
	}
	if len(u.Username) < MinUsernameLength {
		return fmt.Errorf("username must be at least %d characters long", MinUsernameLength)
	}
	if len(u.Username) > MaxUsernameLength {
		return fmt.Errorf("username must be at most %d characters long", MaxUsernameLength)
	}
	if !re.MatchString(u.Username) {
		return fmt.Errorf("username can contain only latin letters, digits, and characters: ! _ -")
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
		return fmt.Errorf("password cannot be empty")
	}
	if len(u.Password) < MinPasswordLength {
		return fmt.Errorf("password must be at least %d characters", MinPasswordLength)
	}
	if len(u.Password) > MaxPasswordLength {
		return fmt.Errorf("password must be at most %d characters", MaxPasswordLength)
	}
	if !re.MatchString(u.Password) {
		return fmt.Errorf("password can contain only latin letters, digits, and characters: ! _ @ .")
	}

	return nil
}
