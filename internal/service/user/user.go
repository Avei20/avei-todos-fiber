package user

import (
	"avei-todos-fiber/internal/entity"
	"context"

	"github.com/google/uuid"
)

func (s *ServiceImpl) Create(
	ctx context.Context, 
	body *CreateBody,
	) (*CreateResponse, error) {
		var (
			user *entity.User
		)

		userRaw := entity.User{
			Id: uuid.NewString(),
			Username: body.Username,
			Email: body.Email,
			Password: body.Password,
		}

		user, err := s.repo.Create(ctx, &userRaw)

		if err != nil {
			return nil, err
		}

		return &CreateResponse{
			Success: true,
			Message: "Success",
			Data: *user,
		}, nil
	}