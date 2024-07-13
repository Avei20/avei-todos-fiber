package handler

import (
	"avei-todos-fiber/internal/handler/project"
	"avei-todos-fiber/internal/handler/todo"
	"avei-todos-fiber/internal/handler/user"
)

func NewHandlers(
	todoHandler todo.Handler,
	projectHandler project.Handler,
	userHandler user.Handler,
) Handlers {
	return Handlers{
		TodoHandler:    todoHandler,
		ProjectHandler: projectHandler,
		UserHandler: userHandler,
	}
}
