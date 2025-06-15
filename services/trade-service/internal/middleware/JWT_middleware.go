package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("your-shared-secret") // You should use env var

func JWTMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization") //Get data from auth in request
	if tokenStr == "" || len(tokenStr) < 8 {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing token or doesn't have Bearer")
	}

	tokenStr = tokenStr[7:]
	//check signing method and decoding!
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unmatch signing method")
		}
		return jwtSecret, nil
	})

	//Check if an error decoding and is it actually valid
	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	//get claims from token
	claims, ok := token.Claims.(jwt.MapClaims) //ok is bool flag, usually return when mapping

	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized,  "Invalid claims")
	}

	//Get user_id from claims
	userIDStr, ok := claims["user_id"].(string)

	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing user_id")
	}

	//Parse string to uuid object
	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID")
	}

	c.Locals("userID", userID)
	return c.Next()
}
