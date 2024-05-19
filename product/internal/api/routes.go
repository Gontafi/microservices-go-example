package api

import (
	"github.com/gofiber/fiber/v2"
)

func (rest *REST) SetupRoutes(router *fiber.App) {
	api := router.Group("/api")

	api.Get("/products/:id", rest.GetProductByID)
	api.Post("/products", VerifyTokenMiddleware, rest.CreateProduct)
	api.Put("/products/:id", VerifyTokenMiddleware, rest.UpdateProductByID)
	api.Delete("/products/:id", VerifyTokenMiddleware, rest.DeleteProductByID)

	api.Post("/products/:id/reviews", VerifyTokenMiddleware, rest.CreateReview)
	api.Delete("/products/:id/reviews/:review_id", VerifyTokenMiddleware, rest.DeleteReviewByID)
}
