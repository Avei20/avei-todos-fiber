package user

import (
	"avei-todos-fiber/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) GetByUsername(
	ctx context.Context,
	username string,
) (*entity.User, error) {
	query := "SELECT * FROM users WHERE username = $1 LIMIT 1"

	rows, err := r.db.Query(ctx, query, username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])

	if err != nil {
		return nil, err
	}

	return &user, nil
}
