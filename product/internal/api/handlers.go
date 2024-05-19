package api

import (
	"github.com/gofiber/fiber/v2"
	"product/internal/models"
	"strconv"
)

func (rest *REST) CreateProduct(c *fiber.Ctx) error {
	var productBody struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
	}

	err := c.BodyParser(&productBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newProduct := models.Product{
		Name:        productBody.Name,
		Description: productBody.Description,
		Price:       productBody.Price,
		Category:    productBody.Category,
	}

	productID, err := rest.service.CreateProduct(c.Context(), newProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "products created", "product_id": productID})
}

func (rest *REST) GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")
	parsedProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid products ID"})
	}

	product, err := rest.service.GetProductByID(c.Context(), parsedProductID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (rest *REST) UpdateProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")
	parsedProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid products ID"})
	}

	var productBody struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
	}

	err = c.BodyParser(&productBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedProduct := models.Product{
		ID:          parsedProductID,
		Name:        productBody.Name,
		Description: productBody.Description,
		Price:       productBody.Price,
		Category:    productBody.Category,
	}

	err = rest.service.UpdateProductByID(c.Context(), updatedProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "products updated"})
}

func (rest *REST) DeleteProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")
	parsedProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid products ID"})
	}

	err = rest.service.DeleteProductByID(c.Context(), parsedProductID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "products deleted"})
}

func (rest *REST) CreateReview(c *fiber.Ctx) error {
	productID := c.Params("id")
	parsedProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid products ID"})
	}

	var reviewBody struct {
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}

	userID, ok := c.Locals("sub").(int64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "can not parse sub from fiber context"})
	}

	err = c.BodyParser(&reviewBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newReview := models.Review{
		ProductID: parsedProductID,
		UserID:    userID,
		Rating:    reviewBody.Rating,
		Comment:   reviewBody.Comment,
	}

	reviewID, err := rest.service.CreateReview(c.Context(), newReview)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "review created", "review_id": reviewID})
}

func (rest *REST) DeleteReviewByID(c *fiber.Ctx) error {
	reviewID := c.Params("review_id")
	parsedReviewID, err := strconv.ParseInt(reviewID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid review ID"})
	}

	err = rest.service.DeleteReviewByID(c.Context(), parsedReviewID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "review deleted"})
}
