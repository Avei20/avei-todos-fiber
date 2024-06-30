package handler

import (
	"avei-todos-fiber/internal/handler/project"
	"avei-todos-fiber/internal/handler/todo"
)

type (
	Handlers struct {
		TodoHandler    todo.Handler
		ProjectHandler project.Handler
	}
)
