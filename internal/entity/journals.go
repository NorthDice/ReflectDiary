package entity

import (
	"fmt"
	"time"
)

// Journal represents a personal journal that belongs to a specific user.
// It contains a title, a reference to the user who owns it, and a list of journal entries.
type Journal struct {
	ID             int       `json:"-"`
	Title          string    `json:"title"`
	UserID         int       `json:"user_id"`
	JournalEntries []Entries `json:"journal_entries"`
	CreatedAt      time.Time `json:"created_at"`
}

// ValidateTitle checks if the journal's title is not empty.
// Returns true if the title is valid, otherwise returns false and an error.
func (j *Journal) ValidateTitle() (bool, error) {
	if j.Title == IsEmptyString {
		return false, fmt.Errorf("title is required")
	}

	return true, nil
}
