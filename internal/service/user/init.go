package user

import "avei-todos-fiber/internal/repository/database/user"

func NewService(repo user.Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}
