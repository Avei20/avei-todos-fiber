package handler

import "avei-todos-fiber/internal/handler/todo"

type (
	Handler interface {
	}

	HandlerImpl struct {
		todoHandler todo.Handler
	}
)
