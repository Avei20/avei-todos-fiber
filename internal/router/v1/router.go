package v1

import (
	"avei-todos-fiber/internal/handler"
	"avei-todos-fiber/internal/router/v1/todo"

	"github.com/gofiber/fiber/v2"
)

func InitV1Router(h *handler.Handler) *fiber.App {
	v1Router := fiber.New()

	v1Router.Mount("/todo", todo.InitRouter(h.todoHandler))

	return v1Router
}
