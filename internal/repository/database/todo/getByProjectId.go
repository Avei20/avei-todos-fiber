package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) GetByProjectId(
	ctx context.Context,
	projectId string,
) ([]entity.Todo, error) {

	query := "SELECT * FROM todos WHERE project_id = $1"

	rows, err := r.db.Query(ctx, query, projectId)

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
