package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"
)

func (s *ServiceImpl) GetAll(ctx context.Context) ([]entity.Todo, error) {
	var (
		todos []entity.Todo
	)

	todos, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}
