package project

import (
	"avei-todos-fiber/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(h *handler.Handlers) *fiber.App {
	router := fiber.New()

	router.Get("/", h.ProjectHandler.Get)
	router.Post("/", h.ProjectHandler.Create)
	return router
}
