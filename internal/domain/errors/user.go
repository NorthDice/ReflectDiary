package domainerr

import (
	apperr "github.com/NorthDice/ReflectDiary/pkg/errors"
)

var (
	ErrUserNotFound = apperr.NewNotFoundError("User")

	ErrUserAlreadyExists     = apperr.NewConflictError("User with this email already exists")
	ErrUsernameAlreadyExists = apperr.NewConflictError("Username already taken")
	ErrUsernameIsRequired    = apperr.NewConflictError("Username is required")
	ErrEmailIsRequired       = apperr.NewConflictError("Email is required")
	ErrInvalidEmail          = apperr.NewValidationError("Invalid email format", nil)
	ErrInvalidUsername       = apperr.NewValidationError("Invalid username", map[string]string{
		"min_length": "3",
		"max_length": "50",
		"pattern":    "alphanumeric characters and underscores only",
	})
	ErrInvalidUsernameMinConstraint = apperr.NewValidationError("Invalid username min constraint", map[string]string{
		"min_length": "3",
	})
	ErrInvalidUsernameMaxConstraint = apperr.NewValidationError("Invalid username max constraint", map[string]string{
		"max_length": "50",
	})
	ErrInvalidUsernameConstraint = apperr.NewValidationError("Invalid username constraint", map[string]string{
		"pattern": "alphanumeric characters and underscores only",
	})
	ErrPasswordIsRequired           = apperr.NewConflictError("Password is required")
	ErrInvalidPasswordMinConstraint = apperr.NewValidationError("Password min constraint", map[string]string{
		"min_length": "6",
	})
	ErrInvalidPasswordMaxConstraint = apperr.NewValidationError("Password max constraint", map[string]string{
		"max_length": "20",
	})
	ErrInvalidPasswordConstraint = apperr.NewValidationError("Password constraint", map[string]string{
		"pattern": "alphanumeric characters and underscores only",
	})
)
