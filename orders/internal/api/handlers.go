package api

import (
	"github.com/gofiber/fiber/v2"
	"orders/internal/models"
	"strconv"
)

func (rest *REST) CreateOrder(c *fiber.Ctx) error {
	var orderBody struct {
		Status      string             `json:"status"`
		TotalAmount float64            `json:"total_amount"`
		Items       []models.OrderItem `json:"items"`
	}

	err := c.BodyParser(&orderBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID, ok := c.Locals("sub").(int64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "can not parse sub from fiber context"})
	}

	newOrder := models.Order{
		UserID:      userID,
		Status:      orderBody.Status,
		TotalAmount: orderBody.TotalAmount,
	}

	orderID, err := rest.service.CreateOrder(c.Context(), newOrder, orderBody.Items)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order created", "order_id": orderID})
}

func (rest *REST) GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")
	parsedOrderID, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid order ID"})
	}

	order, err := rest.service.GetOrderByID(c.Context(), parsedOrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (rest *REST) GetOrderItemsByOrderID(c *fiber.Ctx) error {
	orderID := c.Params("id")
	parsedOrderID, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid order ID"})
	}

	items, err := rest.service.GetOrderItemsByOrderID(c.Context(), parsedOrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(items)
}

func (rest *REST) UpdateOrderStatus(c *fiber.Ctx) error {
	orderID := c.Params("id")
	parsedOrderID, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid order ID"})
	}

	var orderBody struct {
		Status string `json:"status"`
	}

	err = c.BodyParser(&orderBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = rest.service.UpdateOrderStatus(c.Context(), parsedOrderID, orderBody.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order status updated"})
}

func (rest *REST) DeleteOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")
	parsedOrderID, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid order ID"})
	}

	err = rest.service.DeleteOrderByID(c.Context(), parsedOrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order deleted"})
}
