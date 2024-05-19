package api

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"orders/internal/service"
)

type REST struct {
	Router  *fiber.App
	service *service.Service
}

func NewRESTAPI(service *service.Service) *REST {
	rest := REST{}

	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONDecoder:           sonic.Unmarshal,
		JSONEncoder:           sonic.Marshal,
	})

	router.Use(helmet.New(helmet.ConfigDefault))

	router.Use(recover.New(recover.Config{
		Next:              nil,
		EnableStackTrace:  true,
		StackTraceHandler: recover.ConfigDefault.StackTraceHandler,
	}))

	rest.SetupRoutes(router)

	rest.Router = router
	rest.service = service

	return &rest
}
