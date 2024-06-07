package project

import "avei-todos-fiber/internal/repository/neon"

type (
	Repository interface {}

	RepositoryImpl struct {
		db neon.DB
	}
)
