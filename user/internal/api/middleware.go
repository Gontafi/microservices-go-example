package api

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"user/pkg/utils"
)

func JWTVerifyMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing Authorization header"})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization header format"})
	}

	token := parts[1]

	userID, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Token"})
	}

	c.Locals("sub", userID)

	return c.Next()
}
