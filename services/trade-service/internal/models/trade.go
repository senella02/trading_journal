package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Trade struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID     uuid.UUID `gorm:"index" json:"user_id"`
	Symbol     string    `json:"symbol"`
	EntryPrice float64   `json:"entry_price"`
	ExitPrice  float64   `json:"exit_price"`
	LotSize    float64   `json:"lot_size"`
	EntryTime  time.Time `json:"entry_time"`
	ExitTime   time.Time `json:"exit_time"`
	Status     string    `json:"status"`
	Position   string    `json:"position"`
	ImageURL   string    `json:"image_url"`

	PlaybookID uuid.UUID `gorm:"index" json:"playbook_id"`
	Playbook   *Playbook `gorm:"foreignKey:PlaybookID" json:"playbook,omitempty"`

	SetupID uuid.UUID `gorm:"index" json:"setup_id"`
	Setup   *Setup    `gorm:"foreignKey:SetupID" json:"setup,omitempty"`

	Mistakes []Mistake `gorm:"many2many:trade_mistakes;" json:"mistakes"`
	Rules    []Rule    `gorm:"many2many:trade_rules;" json:"rules"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Auto-generate UUID before create
func (t *Trade) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
