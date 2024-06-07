package todo

import "github.com/gofiber/fiber/v2"

func (h *HandlerImpl) Create(c *fiber.Ctx) error {
	return nil
}

func (h *HandlerImpl) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Get todo",
	})
}
