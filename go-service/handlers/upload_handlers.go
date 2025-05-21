package handlers

import (
	
	"os"
	"path/filepath" //Can extract extention of file, create file path to save image
	"strconv"
	"fmt"
	"time"


	"github.com/gofiber/fiber/v2"
	"go-service/utils"
)

type UploadResponse struct {
	ImageURL string `json:"image_url"`
	Filename string `json:"filename"`
}

func UploadHandler(c *fiber.Ctx) error {
	userId := c.Params("userId")
	// Check if the userId is provided
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "userID is required",
		})
	}

	// Check if the file is provided
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file is required",
		})
	}

	// Check if the file is an image
	ext := filepath.Ext(fileHeader.Filename)
	if !utils.IsValidImageExtension(ext) { //not valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid file type",
		})
	}

	//Check size
	maxSizeMB, _ := strconv.Atoi(os.Getenv("MAX_FILE_SIZE"))
	if fileHeader.Size > int64(maxSizeMB*1024*1024) {
		return fiber.NewError(fiber.StatusRequestEntityTooLarge, "File size exceeds limit")
	}

	uniqueName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	userFolder := filepath.Join(os.Getenv("IMAGE_PATH"), userId) //Image path = images -> relative path to working directory(go-service)

	// Create user folder if it doesn't exist
	if err := os.MkdirAll(userFolder, os.ModePerm); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create user folder")
	}

	savePath := filepath.Join(userFolder, uniqueName)
	//Built in fiber function to save file from header
	//Implicitly calling fileHeader.open()
	if err := c.SaveFile(fileHeader, savePath); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not save file")
	}

	return c.JSON(fiber.Map{ //upload is our struct for type safety
		"ImageURL": fmt.Sprintf("/images/%s/%s", userId, uniqueName), //Frontend will access with this image URL!
		"Filename": uniqueName,
	})
}
