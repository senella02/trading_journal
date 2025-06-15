package repository

import (
	"trade-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TradeRepo interface {
	Create(trade *models.Trade) error
	GetByID(id, userID uuid.UUID) (*models.Trade, error) //return pointer to object to prevent redundant copy!
	ListByUser(userID uuid.UUID) ([]models.Trade, error)
	Update(trade *models.Trade, userID uuid.UUID) error
	DeleteByID(id, userID uuid.UUID) error
}

type tradeRepository struct {
	db *gorm.DB
}

func NewTradeRepository(db *gorm.DB) TradeRepo {
	return &tradeRepository{db}
}

// Create inserts a new trade into the database.
func (r *tradeRepository) Create(trade *models.Trade) error {
	return r.db.Create(trade).Error
}

// Delete removes a trade by its ID and userID.
func (r *tradeRepository) DeleteByID(id, userID uuid.UUID) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Trade{}).Error
}

func (r *tradeRepository) GetByID(id, userID uuid.UUID) (*models.Trade, error) {
	t := new(models.Trade) //t is a pointer
	if err := r.db.Preload("Playbook").
		Preload("Setup").
		Preload("Mistake").
		Preload("Rules").
		First(t, "id = ? AND user_id", id, userID).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (r *tradeRepository) ListByUser(userID uuid.UUID) ([]models.Trade, error) {
	var trades []models.Trade
	result := r.db.Find(&trades, "userID = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return trades, nil
}

func (r *tradeRepository) Update(trade *models.Trade, userID uuid.UUID) error {
	var existing models.Trade
	if err := r.db.First(&existing, "id = ? AND user_id = ?", trade.ID, userID).Error; err != nil {
		return err
	}
	return r.db.Save(trade).Error
}
