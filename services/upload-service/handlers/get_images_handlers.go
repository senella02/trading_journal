package handlers

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
)

func GetOneImageHandler(c *fiber.Ctx) error {
	//Get Params
	userId := c.Params("userId")
	fileName := c.Params("fileName")
	//Check if empty
	if userId == "" || fileName == "" {
		return fiber.ErrBadRequest
	}

	//Create file path
	imagePath := filepath.Join(os.Getenv("IMAGE_PATH"), userId, fileName)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return fiber.ErrNotFound
	}

	//Return using send file from fiber
	return c.SendFile(imagePath)

}

func GetImagesHandler(c *fiber.Ctx) error {
	userId := c.Params("userId")
	if userId == "" {
		return fiber.ErrBadRequest
	}
	//Get user folder with relative path
	userFolder := filepath.Join(os.Getenv("IMAGE_PATH"), userId)

	//Read dir return all files into files variable
	files, err := os.ReadDir(userFolder)
	if err != nil {
		return c.JSON([]UploadResponse{})
	}

	var images []UploadResponse //UploadResponse is our return struct for json image response!

	for _, file := range files {
		if file.IsDir() { //files read both file and folder so the thing we read might be a directory
			continue
		}
		//Get name
		name := file.Name()
		
		//Append to response
		images = append(images, UploadResponse{
			ImageURL: fmt.Sprintf("/images/%s/%s", userId, name),
			Filename: name,
		})
	}

	return c.JSON(images) //images = array of struct that has JSON tag
}