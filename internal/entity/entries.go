package entity

import (
	"fmt"
	"strings"
	"time"
)

// Entries represents a single journal entry belonging to a specific journal.
// Each entry has a name (title), content, and timestamps for creation and last update.
type Entries struct {
	ID        int       `json:"-"`
	JournalID int       `json:"-"`
	EntryName string    `json:"entry_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ValidateEntryName checks if the entry name is not empty or whitespace only.
// Returns true if valid, otherwise returns false and an error.
func (e *Entries) ValidateEntryName() (bool, error) {
	if strings.TrimSpace(e.EntryName) == IsEmptyString {
		return false, fmt.Errorf("entry name is required")
	}

	return true, nil
}
