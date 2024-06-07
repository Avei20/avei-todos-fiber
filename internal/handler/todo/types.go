package todo

import (
	"avei-todos-fiber/internal/service/todo"

	"github.com/gofiber/fiber/v2"
)

type (
	HandlerImpl struct {
		todoService todo.Service
	}

	Handler interface {
		RegisterRoutes(router fiber.Router)
	}
)
