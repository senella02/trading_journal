package routes

import (
	"go-service/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//Upload and get images router
	app.Post("/images/:userId", handlers.UploadHandler)   //fiber pass context into handler
	app.Get("/images/:userId", handlers.GetImagesHandler) //Return URL to get one image in form of images/:userId/:filename
	app.Get("/images/:userId/:fileName", handlers.GetOneImageHandler)
}
