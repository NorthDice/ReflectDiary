package postgres

import (
	"context"
	"github.com/NorthDice/ReflectDiary/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserPostgresRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (u *UserRepository) Save(ctx context.Context, user *entity.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var id int

	err := u.pool.QueryRow(ctx, InsertUserIntoUsersTableQuery, user.ID, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// FindById retrieves a user entity by its unique ID.
// Returns the user if found, or an error otherwise.
func (u *UserRepository) FindById(ctx context.Context, id int) (*entity.User, error) {
	return nil, nil
}

// FindByEmail retrieves a user entity by its email address.
// Returns the user if found, or an error otherwise.
func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

// Update modifies an existing user entity.
// Returns an error if the update fails.
func (u *UserRepository) Update(ctx context.Context, user *entity.User) error {
	return nil
}

// Delete removes a user entity by its ID.
// Returns an error if deletion fails.
func (u *UserRepository) Delete(ctx context.Context, id int) error {
	return nil
}
