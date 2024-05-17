package todo

import (
	"avei-todos-fiber/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(h *handler.Handler) *fiber.App {
	router := fiber.New()

	return router
}
