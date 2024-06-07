package handler

import "avei-todos-fiber/internal/handler/todo"

type (
	Handlers struct {
		todoHandler *todo.Handler
	}
)
