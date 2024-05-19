package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"orders/config"
	authpb "orders/gen/go/auth"
)

func VerifyTokenMiddleware(c *fiber.Ctx) error {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", config.AppConfig.AuthAppHost, config.AppConfig.AuthGRPCPort),
		grpc.WithInsecure())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Authentication Failed",
			"error":   err.Error(),
		})
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	response, err := client.VerifyToken(c.Context(), &authpb.VerifyTokenRequest{Token: c.Get("Authorization")})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"message": "Authentication Failed, please insert walid token",
				"error":   err.Error(),
			})
	}

	if !response.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": response.Message})
	}

	c.Locals("sub", response.UserID)

	return c.Next()
}
