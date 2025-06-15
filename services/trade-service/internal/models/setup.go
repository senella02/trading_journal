package models

import (
	"time"

	"github.com/google/uuid"
)

// Setup entity
type Setup struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	PlaybookID  uuid.UUID `gorm:"index" json:"playbook_id"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
	Rules       []Rule    `gorm:"foreignKey:SetupID" json:"rules"`
	Trades      []Trade   `gorm:"foreignKey:SetupID" json:"trades"` // Trade has fk setup_id
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
