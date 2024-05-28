package app

import (
	"avei-todos-fiber/internal/handler/todo"
)

type Server struct {
	todoHandler todo.Handler
}
