package todo

import (
	"avei-todos-fiber/internal/service/todo"

	"github.com/gofiber/fiber"
)

type (
	HandlerImpl struct {
		todoService todo.Service
	}

	Handler interface {
		RegisterRoutes(router fiber.Router)
	}
)
