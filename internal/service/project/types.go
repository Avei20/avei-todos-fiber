package project

import (
	"avei-todos-fiber/internal/entity"
	"avei-todos-fiber/internal/repository/database/project"
	"context"
)

type (
	Service interface {
		GetAll(ctx context.Context) (*GetAllResponse, error)
		Create(ctx context.Context, body *CreateBody) (*CreateResponse, error)
	}

	ServiceImpl struct {
		repo project.Repository
	}

	GetAllResponse struct {
		Success bool             `json:"success"`
		Message string           `json:"message"`
		Data    []entity.Project `json:"data"`
	}

	CreateResponse struct {
		Success bool           `json:"success"`
		Message string         `json:"message"`
		Data    entity.Project `json:"data"`
	}

	CreateBody struct {
		Code        string  `json:"code"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}
)
