package repository

import (
	"trade-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type PlaybookRepo interface {
    Create(playbook *models.Playbook) error
    GetByID(id uuid.UUID) (*models.Playbook, error)
    ListByUser(userID uuid.UUID) ([]models.Playbook, error)
    Update(playbook *models.Playbook) error
    DeleteByID(id uuid.UUID) error
}

type playbookRepository struct {
    db *gorm.DB
}

func NewPlaybookRepository(db *gorm.DB) PlaybookRepo {
    return &playbookRepository{db}
}

func (r *playbookRepository) Create(playbook *models.Playbook) error {
    return r.db.Create(playbook).Error
}

func (r *playbookRepository) GetByID(id uuid.UUID) (*models.Playbook, error) {
    var pb models.Playbook
    if err := r.db.Preload("Setups").Preload("Mistakes").First(&pb, "id = ?", id).Error; err != nil{
			return nil, err
		}
    return &pb, nil
}

func (r *playbookRepository) ListByUser(userID uuid.UUID) ([]models.Playbook, error) {
    var playbooks []models.Playbook
    if err := r.db.Find(&playbooks, "userID = ?", userID).Error; err != nil {
		return nil, err
	}
    return playbooks, nil
}

func (r *playbookRepository) Update(playbook *models.Playbook) error {
    if err := r.db.Save(playbook).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *playbookRepository) DeleteByID(id uuid.UUID) error {
    if err := r.db.Where("id = ?", id).Delete(&models.Playbook{}).Error;
		err != nil {
		return err
	}
	return nil
}