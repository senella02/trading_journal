package models

import (
	"time"

	"github.com/google/uuid"
)

// Playbook entity
type Playbook struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID      uuid.UUID `gorm:"index" json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Emoji       string    `json:"emoji"`
	Setups      []Setup   `gorm:"foreignKey:PlaybookID" json:"setups"`
	Mistakes    []Mistake `gorm:"many2many:playbook_mistakes;" json:"mistakes"`
	Trades      []Trade   `gorm:"foreignKey:PlaybookID" json:"trades"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
