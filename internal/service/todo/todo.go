package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"
)

func (s *ServiceImpl) GetAll(ctx context.Context) ([]entity.Todo, error) {
	return s.repo.GetAll(ctx)
}
