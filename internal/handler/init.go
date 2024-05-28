package handler

import "avei-todos-fiber/internal/handler/todo"

func NewHandler(
	todoHandler todo.Handler,
) Handler {
	return &HandlerImpl{
		todoHandler: todoHandler,
	}
}
