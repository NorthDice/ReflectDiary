package postgres

import (
	"github.com/NorthDice/ReflectDiary/internal/usecase/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	User interfaces.UserRepository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User: NewUserPostgresRepository(db),
	}
}
