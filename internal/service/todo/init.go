package todo

import "avei-todos-fiber/internal/repository/todo"

func NewService(repo todo.Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}
