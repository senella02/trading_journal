package repository

import (
	"trade-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SetupRepo interface {
	BulkCreate(setups []models.Setup) error

	GetByID(id uuid.UUID) (*models.Setup, error)

	ListByPlaybookID(playbookID uuid.UUID) ([]models.Setup, error)
}

type setupRepository struct {
	db *gorm.DB
}

func NewSetupRepository(db *gorm.DB) SetupRepo {
	return &setupRepository{db}
}

func (r *setupRepository) BulkCreate(setups []models.Setup) error {
	return r.db.Create(&setups).Error
}

func (r *setupRepository) GetByID(id uuid.UUID) (*models.Setup, error) {
	var setup models.Setup
	if err := r.db.First(&setup, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &setup, nil
}

func (r *setupRepository) ListByPlaybookID(playbookID uuid.UUID) ([]models.Setup, error) {
	var setups []models.Setup
	if err := r.db.Where("playbook_id = ?", playbookID).Find(&setups).Error; err != nil {
		return nil, err
	}
	return setups, nil
}
