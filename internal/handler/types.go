package handler

import "avei-todos-fiber/internal/handler/todo"

type (
	Handlers interface {
	}

	HandlerImpl struct {
		todoHandler todo.Handler
	}
)
