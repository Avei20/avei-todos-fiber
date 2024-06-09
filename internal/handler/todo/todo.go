package todo

import (
	"avei-todos-fiber/internal/entity"
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
//	@Success		201	{object}	interface{}
//	@Router			/v1/todos [post]
func (h *HandlerImpl) Create(c *fiber.Ctx) error {
	return nil
}

// GetTodo godoc
//
//	@Summary		Get todo
//	@Description	Get todo
//	@Tags			Todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interface{}
//	@Response		200	{object}	interface{}
//	@Response		404	{object}	interface{}
//	@Response		500	{object}	interface{}
//	@Router			/v1/todos/ [get]
func (h *HandlerImpl) Get(c *fiber.Ctx) error {
	var (
		todos []entity.Todo
		err   error
	)

	// IF param null

	todos, err = h.todoService.GetAll(c.Context())

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengambil data todos!",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Ini data semua todosnya ya bang!",
		"data":    todos,
	})
}
