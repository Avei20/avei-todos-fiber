package todo

import (
	"avei-todos-fiber/internal/entity"
	"avei-todos-fiber/internal/repository/neon"
	"context"
)

type (
	Repository interface {
		GetAll(ctx context.Context) ([]entity.Todo, error)
		Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error)
	}

	RepositoryImpl struct {
		db neon.DB
	}
)
