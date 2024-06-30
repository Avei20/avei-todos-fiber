package todo

import (
	"avei-todos-fiber/internal/entity"
	"avei-todos-fiber/internal/repository/database/todo"
	"context"
)

type (
	Service interface {
		GetAll(ctx context.Context) (*GetAllResponse, error)
		Create(ctx context.Context, body *CreateBody) (*CreateResponse, error)
	}
	ServiceImpl struct {
		repo todo.Repository
	}

	GetAllResponse struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []entity.Todo `json:"data"`
	}

	CreateResponse struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    entity.Todo `json:"data"`
	}

	CreateBody struct {
		Name      string  `json:"name"`
		ProjectId *string `json:"project_id"`
		ParentId  *string `json:"parent_id"`
	}
)
