package project

import "avei-todos-fiber/internal/repository/database/project"

func NewService(repo project.Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}
