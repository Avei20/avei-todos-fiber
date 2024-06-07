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
		Create(c *fiber.Ctx) error
		Get(c *fiber.Ctx) error
	}
)
