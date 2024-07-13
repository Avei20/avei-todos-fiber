package user

import (
	"avei-todos-fiber/internal/entity"
	"avei-todos-fiber/internal/repository/neon"
	"context"
)

type (
	Repository interface {
		GetAll(ctx context.Context) ([]entity.User, error)
		Create(ctx context.Context, user *entity.User) (*entity.User, error)
	}

	RepositoryImpl struct {
		db neon.DB
	}
)
