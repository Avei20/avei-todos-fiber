package handler

import "avei-todos-fiber/internal/handler/todo"

func NewHandlers(
	todoHandler todo.Handler,
) Handlers {
	return Handlers{
		TodoHandler: todoHandler,
	}
}
