package routes

import (
    "github.com/gofiber/fiber/v2"
    "trade-service/internal/handlers"
    "trade-service/internal/repository"
    "gorm.io/gorm"
)

//gorm.DB = connection object that shared in fiber
func SetupTradeRoutes(app *fiber.App, db *gorm.DB) {
    tradeRepo := repository.NewTradeRepository(db)
    tradeHandler := handlers.NewTradeHandler(tradeRepo)

    tradeGroup := app.Group("/trades")
    tradeGroup.Post("/", tradeHandler.Create)
    tradeGroup.Get("/", tradeHandler.ListByUser)
    tradeGroup.Get("/:id", tradeHandler.GetByID)
    tradeGroup.Put("/:id", tradeHandler.Update)
    tradeGroup.Delete("/:id", tradeHandler.Delete)

    playbookRepo := repository.NewPlaybookRepository(db)
    playbookHandler := handlers.NewPlaybookHandler(playbookRepo)
    playbookGroup := app.Group("/playbooks")

    playbookGroup.Get("/:id", playbookHandler.GetByID) //get playbook by playbook id
    playbookGroup.Get("/", playbookHandler.ListByUser) //get user id from middleware
    playbookGroup.Post("/", playbookHandler.Create)
    playbookGroup.Put("/:id", tradeHandler.Update)
    playbookGroup.Delete("/:id", tradeHandler.Delete)
}


