package project

import (
	"avei-todos-fiber/internal/entity"
	"avei-todos-fiber/internal/repository/neon"
	"context"
)

type (
	Repository interface {
		GetAll(ctx context.Context) ([]entity.Project, error)
		Create(ctx context.Context, project *entity.Project) (*entity.Project, error)
	}

	RepositoryImpl struct {
		db neon.DB
	}
)
