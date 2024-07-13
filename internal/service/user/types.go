package user

import (
	"avei-todos-fiber/internal/entity"
	"avei-todos-fiber/internal/repository/database/user"
	"context"
)

type (
	Service interface {
		// GetByUsername(ctx context.Context, username string) (*GetByUsernameResponse, error)
		Create(ctx context.Context, body *CreateBody) (*CreateResponse, error)
	}

	ServiceImpl struct {
		repo user.Repository
	}

	CreateBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	GetByUsernameResponse struct {
		Success bool         `json:"success"`
		Message string       `json:"message"`
		Data    *entity.User `json:"data"`
	}

	CreateResponse struct {
		Success bool         `json:"success"`
		Message string       `json:"message"`
		Data    entity.User `json:"data"`
	}
)
