package todo

import (
	"avei-todos-fiber/internal/service/todo"
)

type (
	HandlerImpl struct {
		todoService todo.Service
	}

	Handler interface {
	}
)
