package project

import (
	"avei-todos-fiber/internal/entity"
	"context"
	"log"

	"github.com/google/uuid"
)

func (s *ServiceImpl) GetAll(ctx context.Context) (*GetAllResponse, error) {
	projects, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &GetAllResponse{
		Success: true,
		Message: "Projects retrieved successfully",
		Data:    projects,
	}, nil
}

func (s *ServiceImpl) Create(ctx context.Context, body *CreateBody) (*CreateResponse, error) {
	var (
		project *entity.Project
	)

	projectRaw := entity.Project{
		Name:        body.Name,
		Code:        body.Code,
		Description: *body.Description,
		Id:          uuid.NewString(),
		Deleted:     false,
	}

	project, err := s.repo.Create(ctx, &projectRaw)

	if err != nil {
		log.Println("Vah Error Bwang")
		return nil, err
	}

	return &CreateResponse{
		Success: true,
		Message: "Project created successfully",
		Data:    *project,
	}, nil

}
