package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "trade-service/internal/models"
    "trade-service/internal/repository"
)

type TradeHandler struct {
    repo repository.TradeRepo
}

//init
func NewTradeHandler(repo repository.TradeRepo) *TradeHandler {
    return &TradeHandler{repo}
}

// POST /trades
func (h *TradeHandler) Create(c *fiber.Ctx) error {
	trade := new(models.Trade)
	if err := c.BodyParser(trade); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Extract user ID (assumes middleware adds this)
	userID := c.Locals("userID").(uuid.UUID)
	trade.UserID = userID

	if err := h.repo.Create(trade); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(trade)
}

// GET /trades/:id
func (h *TradeHandler) GetByID(c *fiber.Ctx) error {
	tradeID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid trade ID")
	}
	userID := c.Locals("userID").(uuid.UUID)

	trade, err := h.repo.GetByID(tradeID, userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Trade not found")
	}
	return c.JSON(trade)
}

// GET /trades
func (h *TradeHandler) ListByUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	trades, err := h.repo.ListByUser(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(trades)
}

// PUT /trades/:id
func (h *TradeHandler) Update(c *fiber.Ctx) error {
	tradeID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid trade ID")
	}
	userID := c.Locals("userID").(uuid.UUID)

	updated := new(models.Trade)
	if err := c.BodyParser(updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	updated.ID = tradeID
	updated.UserID = userID

	if err := h.repo.Update(updated, userID); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Update failed: "+err.Error())
	}
	return c.JSON(updated)
}

// DELETE /trades/:id
func (h *TradeHandler) Delete(c *fiber.Ctx) error {
	tradeID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid trade ID")
	}
	userID := c.Locals("userID").(uuid.UUID)

	if err := h.repo.DeleteByID(tradeID, userID); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Delete failed: "+err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
