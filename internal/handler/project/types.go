package project

import (
	"avei-todos-fiber/internal/service/project"

	"github.com/gofiber/fiber/v2"
)

type (
	HandlerImpl struct {
		projectService project.Service
	}

	Handler interface {
		Create(c *fiber.Ctx) error
		Get(c *fiber.Ctx) error
	}
)
