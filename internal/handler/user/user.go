package user

import (
	"avei-todos-fiber/internal/service/user"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc

// @Summary 	Create an User
// @Description Create an User
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		body	body		user.CreateBody	true	"Create User"
// @Success	201		{object}	user.CreateResponse
// @Router 		/v1/users [post]
func (h *HandlerImpl) Create(c *fiber.Ctx) error {
	body := new(user.CreateBody)
	err := c.BodyParser(body)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid request body",
				"data":    nil,
			})
	}

	response, err := h.userService.Create(c.Context(), body)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Internal server error",
				"data":    nil,
			})
	}

	return c.JSON(response)
}
