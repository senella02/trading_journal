package routes

import (
	"upload-service/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//Upload and get images router
	api := app.Group("/api/v1")
	images := api.Group("/images")

	images.Post("/images/:userId", handlers.UploadHandler)   //fiber pass context into handler
	images.Get("/images/:userId", handlers.GetImagesHandler) //Return URL to get one image in form of images/:userId/:filename
	images.Get("/images/:userId/:fileName", handlers.GetOneImageHandler)
}
