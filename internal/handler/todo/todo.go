package todo

import (
	"avei-todos-fiber/internal/service/todo"
	"log"

	"github.com/gofiber/fiber/v2"
)

// CreateTodo godoc
//
//	@Summary		Create a todo
//	@Description	Create a todo
//	@Tags			Todos
//	@Accept			json
//	@Produce		json
//	@Param			body body todo.CreateBody true "Create Todo"
//	@Success		201	{object}	todo.CreateResponse
//	@Router			/v1/todos [post]
func (h *HandlerImpl) Create(c *fiber.Ctx) error {
	body := new(todo.CreateBody)
	err := c.BodyParser(body)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal membuat todo!",
			"data":    nil,
		})
	}

	response, err := h.todoService.Create(c.Context(), body)

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

// GetTodo godoc
//
//	@Summary		Get todo
//	@Description	Get todo
//	@Tags			Todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	todo.GetAllResponse
//	@Router			/v1/todos/ [get]
func (h *HandlerImpl) Get(c *fiber.Ctx) error {
	var (
		err error
	)

	// IF param null

	response, err := h.todoService.GetAll(c.Context())

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengambil data todos!",
			"data":    nil,
		})
	}

	return c.JSON(response)
}
