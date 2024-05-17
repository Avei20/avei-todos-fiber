package router

import (
	"avei-todos-fiber/internal/handler"
	v1 "avei-todos-fiber/internal/router/v1"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(h *handler.Handler) *fiber.App {
	router := fiber.New()

	v1Router := v1.InitV1Router(h)
	router.Mount("/v1", v1Router)

	return router
}
