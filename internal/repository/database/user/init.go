package user

import "avei-todos-fiber/internal/repository/neon"

func NewRepository(db neon.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}
