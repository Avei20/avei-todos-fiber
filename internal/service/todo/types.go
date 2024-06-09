package todo

import (
	"avei-todos-fiber/internal/entity"
	"avei-todos-fiber/internal/repository/database/todo"
	"context"
)

type (
	Service interface {
		GetAll(ctx context.Context) ([]entity.Todo, error)
	}
	ServiceImpl struct {
		repo todo.Repository
	}
)
