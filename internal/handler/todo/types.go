package todo

import "avei-todos-fiber/internal/service/todo"

type (
	Handler struct {
		todoService todo.Service
	}
)
