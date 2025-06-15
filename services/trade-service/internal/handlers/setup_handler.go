package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"trade-service/internal/models"
	"trade-service/internal/repository"
)

type SetupHandler struct {
	repo         repository.SetupRepo
	playbookRepo repository.PlaybookRepo
}

func NewSetupHandler(repo repository.SetupRepo, playbookRepo repository.PlaybookRepo) *SetupHandler {
	return &SetupHandler{repo, playbookRepo}
}

// POST /playbooks/:id/setups
func (h *SetupHandler) CreateBatch(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	playbookID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	// verify ownership
	pb, err := h.playbookRepo.GetByID(playbookID)
	if err != nil {
		return fiber.ErrNotFound
	}
	if pb.UserID != userID {
		return fiber.ErrForbidden
	}

	var setups []models.Setup
	if err := c.BodyParser(&setups); err != nil {
		return fiber.ErrBadRequest
	}

	for i := range setups {
		setups[i].PlaybookID = playbookID
	}

	if err := h.repo.BulkCreate(setups); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(setups)
}

// GET /playbooks/:id/setups
func (h *SetupHandler) ListByPlaybook(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	playbookID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	// verify ownership
	pb, err := h.playbookRepo.GetByID(playbookID)
	if err != nil {
		return fiber.ErrNotFound
	}
	if pb.UserID != userID {
		return fiber.ErrForbidden
	}

	setups, err := h.repo.ListByPlaybookID(playbookID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(setups)
}

// GET /setups/:id
func (h *SetupHandler) GetByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	setup, err := h.repo.GetByID(id)
	if err != nil {
		return fiber.ErrNotFound
	}

	// simple ownership check via playbook
	pb, err := h.playbookRepo.GetByID(setup.PlaybookID)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if pb.UserID != userID {
		return fiber.ErrForbidden
	}

	return c.JSON(setup)
}
