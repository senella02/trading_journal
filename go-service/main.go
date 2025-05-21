package main

import (
	"os"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"

	"go-service/config"
	"go-service/routes"
)

func main() {
	//Setup config
	config.LoadEnv()

	//Init server
	app := fiber.New()
	app.Use(logger.New())
	app.Static("/images", "./images")

	//Setup routes
	routes.SetupRoutes(app)

	//Setup Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	app.Listen(":" + port) //:8081
	log.Printf("Go Fiber Server running on http://localhost/:%s\n", port)
	
	
}
