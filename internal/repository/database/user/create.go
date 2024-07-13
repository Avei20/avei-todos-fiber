package user

import (
	"avei-todos-fiber/internal/entity"
	"context"
)

func (r *RepositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `INSERT
	INTO users (id, username, email, password)
	VALUES ($1, $2, $3, $4)`

	rows, err := r.db.Exec(
		ctx,
		query,
		user.Id,
		user.Username,
		user.Email,
		user.Password,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
