package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "trade-service/internal/models"
    "trade-service/internal/repository"
)

type PlaybookHandler struct {
    repo repository.PlaybookRepo
}

//init
func NewPlaybookHandler(repo repository.PlaybookRepo) *PlaybookHandler {
    return &PlaybookHandler{repo}
}

func (h *PlaybookHandler) Create(c *fiber.Ctx) error {
    userID := c.Locals("userID").(uuid.UUID)
    if userID == uuid.Nil {
        return fiber.ErrUnauthorized
    }

    var pb models.Playbook
    if err := c.BodyParser(&pb); err != nil {
        return fiber.ErrBadRequest
    }

    pb.UserID = userID

    if err := h.repo.Create(&pb); err != nil {
        return fiber.ErrInternalServerError
    }

    return c.Status(fiber.StatusCreated).JSON(pb)
}

func (h *PlaybookHandler) ListByUser(c *fiber.Ctx) error {
    userID := c.Locals("userID").(uuid.UUID)
    if userID == uuid.Nil {
        return fiber.ErrUnauthorized
    }

    playbooks, err := h.repo.ListByUser(userID)
    if err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(playbooks)
}

func (h *PlaybookHandler) GetByID(c *fiber.Ctx) error {
    userID := c.Locals("userID").(uuid.UUID)
    if userID == uuid.Nil {
        return fiber.ErrUnauthorized
    }

    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return fiber.ErrBadRequest
    }

    pb, err := h.repo.GetByID(id)
    if err != nil {
        return fiber.ErrNotFound
    }

    if pb.UserID != userID {
        return fiber.ErrForbidden
    }

    return c.JSON(pb)
}

func (h *PlaybookHandler) Update(c *fiber.Ctx) error {
    userID := c.Locals("userID").(uuid.UUID)
    if userID == uuid.Nil {
        return fiber.ErrUnauthorized
    }

    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return fiber.ErrBadRequest
    }

    pb, err := h.repo.GetByID(id)
    if err != nil {
        return fiber.ErrNotFound
    }

    if pb.UserID != userID {
        return fiber.ErrForbidden
    }

    var update models.Playbook
    if err := c.BodyParser(&update); err != nil {
        return fiber.ErrBadRequest
    }

    pb.Name = update.Name
    pb.Description = update.Description
    pb.Emoji = update.Emoji
	pb.Mistakes = update.Mistakes

    if err := h.repo.Update(pb); err != nil {
        return fiber.ErrInternalServerError
    }

    return c.JSON(pb)
}

func (h *PlaybookHandler) Delete(c *fiber.Ctx) error {
     userID := c.Locals("userID").(uuid.UUID)
    if userID == uuid.Nil {
        return fiber.ErrUnauthorized
    }

    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return fiber.ErrBadRequest
    }

    pb, err := h.repo.GetByID(id)
    if err != nil {
        return fiber.ErrNotFound
    }

    if pb.UserID != userID {
        return fiber.ErrForbidden
    }

    if err := h.repo.DeleteByID(id); err != nil {
        return fiber.ErrInternalServerError
    }

    return c.SendStatus(fiber.StatusNoContent)
}
