package handler

import (
	"avei-todos-fiber/internal/handler/project"
	"avei-todos-fiber/internal/handler/todo"
)

func NewHandlers(
	todoHandler todo.Handler,
	projectHandler project.Handler,
) Handlers {
	return Handlers{
		TodoHandler:    todoHandler,
		ProjectHandler: projectHandler,
	}
}
