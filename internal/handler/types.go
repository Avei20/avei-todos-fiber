package handler

import (
	"avei-todos-fiber/internal/handler/project"
	"avei-todos-fiber/internal/handler/todo"
	"avei-todos-fiber/internal/handler/user"
)

type (
	Handlers struct {
		TodoHandler    todo.Handler
		ProjectHandler project.Handler
		UserHandler user.Handler
	}
)
