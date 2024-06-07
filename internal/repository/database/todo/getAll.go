package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) GetAll(ctx context.Context) ([]entity.Todo, error) {

	query := "SELECT * FROM todos"

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Todo])

	if err != nil {
		return nil, err
	}

	return todos, nil
}
