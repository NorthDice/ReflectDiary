package interfaces

import (
	"context"
	"github.com/NorthDice/ReflectDiary/internal/entity"
)

// UserRepository defines methods for managing user data persistence.
type UserRepository interface {
	// Save stores a new user entity in the repository.
	// Returns the new user's ID or an error if saving fails.
	Save(ctx context.Context, user *entity.User) (int, error)

	// FindById retrieves a user entity by its unique ID.
	// Returns the user if found, or an error otherwise.
	FindById(ctx context.Context, id int) (*entity.User, error)

	// FindByEmail retrieves a user entity by its email address.
	// Returns the user if found, or an error otherwise.
	FindByEmail(ctx context.Context, email string) (*entity.User, error)

	// Update modifies an existing user entity.
	// Returns an error if the update fails.
	Update(ctx context.Context, user *entity.User) error

	// Delete removes a user entity by its ID.
	// Returns an error if deletion fails.
	Delete(ctx context.Context, id int) error
}
