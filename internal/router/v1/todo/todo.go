package todo

import (
	"avei-todos-fiber/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(h *handler.Handlers) *fiber.App {
	router := fiber.New()

	router.Get("/", h.TodoHandler.Get)
	router.Post("/", h.TodoHandler.Create)
	return router
}
