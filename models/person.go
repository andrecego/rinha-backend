package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Person struct {
	ID       uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4();not null"`
	Nickname string         `json:"apelido"`
	Name     string         `json:"nome"` // gorm:"not null;unique"
	Birthday Date           `json:"nascimento"`
	Stack    pq.StringArray `json:"stack" gorm:"type:text[]"`
}

// IMPROVEMENT: check nickname uniqueness
func (p Person) Validate() bool {
	// Nickname	has to be at most 32 characters long.
	if len(p.Nickname) == 0 {
		return false
	}
	if len(p.Nickname) > 32 {
		return false
	}

	// Name	has to be at most 100 characters long.
	if len(p.Name) == 0 {
		return false
	}
	if len(p.Name) > 100 {
		return false
	}

	// Stack if present, has to be a list of strings with at most 32 characters each.
	for _, s := range p.Stack {
		if len(s) > 32 {
			return false
		}
	}

	return true
}
