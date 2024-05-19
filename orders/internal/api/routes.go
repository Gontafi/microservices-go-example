package api

import (
	"github.com/gofiber/fiber/v2"
)

func (rest *REST) SetupRoutes(router *fiber.App) {
	api := router.Group("/api")

	api.Get("/orders/:id", rest.GetOrderByID)
	api.Post("/orders", VerifyTokenMiddleware, rest.CreateOrder)
	api.Put("/orders/:id/status", VerifyTokenMiddleware, rest.UpdateOrderStatus)
	api.Delete("/orders/:id", VerifyTokenMiddleware, rest.DeleteOrderByID)
	api.Get("/orders/:id/items", VerifyTokenMiddleware, rest.GetOrderItemsByOrderID)
}
