package user

import (
	"avei-todos-fiber/internal/service/user"

	"github.com/gofiber/fiber/v2"
)

type (
	HandlerImpl struct {
		userService user.Service
	}

	Handler interface {
		Create(c *fiber.Ctx) error
	}
)
