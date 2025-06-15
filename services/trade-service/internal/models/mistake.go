package models

import (
	"time"

	"github.com/google/uuid"
)

// Mistake entity
type Mistake struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;" json:"id"`
	Name        string     `json:"name"`
	Color       string     `json:"color"`
	Description string     `json:"description"`
	Playbooks   []Playbook `gorm:"many2many:playbook_mistakes;" json:"playbooks"`
	Trades      []Trade    `gorm:"many2many:trade_mistakes;" json:"trades"` // Inverse relation
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
