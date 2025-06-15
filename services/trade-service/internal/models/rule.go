package models

import (
	"time"

	"github.com/google/uuid"
)

// Rule entity
type Rule struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
	SetupID     uuid.UUID `gorm:"index" json:"setup_id"`
	Trades      []Trade   `gorm:"many2many:trade_rules;" json:"trades"` // Inverse relation
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
