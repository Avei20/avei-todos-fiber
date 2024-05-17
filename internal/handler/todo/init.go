package todo

import "avei-todos-fiber/internal/service/todo"

func NewHandler(todoService todo.Service) *Handler {
	return &Handler{
		todoService: todoService,
	}
}
