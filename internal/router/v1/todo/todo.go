package todo

import (
	"avei-todos-fiber/internal/handler/todo"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(h *todo.Handler) *fiber.App {
	router := fiber.New()

	return router
}