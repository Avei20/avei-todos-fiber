package todo

import "avei-todos-fiber/internal/repository/neon"

type (
	Repository interface {
	}

	RepositoryImpl struct {
		db neon.DB
	}
)
