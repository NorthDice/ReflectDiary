package entity

import (
	"fmt"
	"strings"
	"time"
)

type Entries struct {
	ID        int       `json:"-"`
	JournalID int       `json:"-"`
	EntryName string    `json:"entry_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Entries) ValidateEntryName() (bool, error) {
	if strings.TrimSpace(e.EntryName) == IsEmptyString {
		return false, fmt.Errorf("entry name is required")
	}

	return true, nil
}
