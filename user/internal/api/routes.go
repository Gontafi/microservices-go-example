package api

import "github.com/gofiber/fiber/v2"

func (rest *REST) SetupRoutes(router *fiber.App) {
	api := router.Group("api/")

	api.Post("/register", rest.RegisterUser)
	api.Post("/login", rest.LoginUser)
	api.Post("/reset-password", rest.ResetPassword)
	api.Get("/profile", JWTVerifyMiddleware, rest.GetUserInfo)
	api.Put("/profile", JWTVerifyMiddleware, rest.ChangeUserProfile)
}
