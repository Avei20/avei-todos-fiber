package v1

import (
	"avei-todos-fiber/internal/handler"
	"avei-todos-fiber/internal/router/v1/project"
	"avei-todos-fiber/internal/router/v1/todo"
	"avei-todos-fiber/internal/router/v1/user"

	"github.com/gofiber/fiber/v2"
)

func InitV1Router(h *handler.Handlers) *fiber.App {
	v1Router := fiber.New()

	v1Router.Mount("/todos", todo.InitRouter(h))
	v1Router.Mount("/projects", project.InitRouter(h))
	v1Router.Mount("/users", user.InitRouter(h))

	return v1Router
}
