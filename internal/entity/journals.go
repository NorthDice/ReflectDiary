package entity

import (
	"fmt"
	"time"
)

type Journal struct {
	ID        int       `json:"-"`
	Title     string    `json:"title"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (j *Journal) ValidateTitle() (bool, error) {
	if j.Title == IsEmptyString {
		return false, fmt.Errorf("title is required")
	}

	return true, nil
}
