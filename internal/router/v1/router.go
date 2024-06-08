package v1

import (
	"avei-todos-fiber/internal/handler"
	"avei-todos-fiber/internal/router/v1/todo"

	"github.com/gofiber/fiber/v2"
)

func InitV1Router(h *handler.Handlers) *fiber.App {
	v1Router := fiber.New()

	v1Router.Mount("/todos", todo.InitRouter(h))

	return v1Router
}
