package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Person struct {
	ID       uuid.UUID      `json:"id,omitempty" gorm:"type:uuid;primary_key;default:uuid_generate_v4();not null"`
	Nickname string         `json:"apelido,omitempty" gorm:"unique"`
	Name     string         `json:"nome,omitempty"` // gorm:"not null;unique"
	Birthday Date           `json:"nascimento,omitempty"`
	Stack    pq.StringArray `json:"stack,omitempty" gorm:"type:text[]"`
}

const (
	MaxNameLength     = 100
	MaxNicknameLength = 32
	MaxStackLength    = 32
)

// IMPROVEMENT: check nickname uniqueness
func (p Person) Validate() bool {
	if len(p.Nickname) == 0 {
		return false
	}

	if len(p.Name) == 0 {
		return false
	}

	if len(p.Name) > MaxNameLength {
		return false
	}

	if len(p.Nickname) > MaxNicknameLength {
		return false
	}

	for _, s := range p.Stack {
		if len(s) > MaxStackLength {
			return false
		}
	}

	return true
}
