package project

import (
	"avei-todos-fiber/internal/service/project"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Create Project godoc
// @Summary Create Project
// @Description Create Project
// @Tags Project
// @Accept json
// @Produce json
// @Param body body project.CreateBody true "Create Project"
// @Success 201 {object} project.CreateResponse
// @Router /v1/projects [post]
func (h *HandlerImpl) Create(c *fiber.Ctx) error {
	body := new(project.CreateBody)
	err := c.BodyParser(body)

	log.Printf("%+v", body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal membuat project!",
			"data":    nil,
		})
	}

	response, err := h.projectService.Create(c.Context(), body)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.JSON(response)
}

// Get Project godoc
// @Summary Get Project
// @Description Get Project
// @Tags Project
// @Accept json
// @Produce json
// @Success 200 {object} project.GetAllResponse
// @Router /v1/projects/ [get]
func (h *HandlerImpl) Get(c *fiber.Ctx) error {
	var (
		err error
	)

	response, err := h.projectService.GetAll(c.Context())

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.JSON(response)
}
