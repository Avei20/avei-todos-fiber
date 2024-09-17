package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"
	"log"

	"github.com/google/uuid"
)

func (s *ServiceImpl) GetAll(ctx context.Context) (*GetAllResponse, error) {
	var (
		todos []entity.Todo
	)

	todos, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return &GetAllResponse{
		Success: true,
		Message: "Success",
		Data:    todos,
	}, nil
}

func (s *ServiceImpl) Create(ctx context.Context, body *CreateBody) (*CreateResponse, error) {
	var (
		todo *entity.Todo
	)

	log.Print("%+v", body)

	todoRaw := entity.Todo{
		Id:        uuid.NewString(),
		Name:      body.Name,
		ProjectId: body.ProjectId,
		ParentId:  body.ParentId,
		Deleted:   false,
		IsDone:    false,
	}

	log.Printf("%+v", todoRaw)

	todo, err := s.repo.Create(ctx, &todoRaw)

	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		Success: true,
		Message: "Success",
		Data:    *todo,
	}, nil
}
