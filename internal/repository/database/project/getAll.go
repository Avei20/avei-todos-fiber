package project

import (
	"avei-todos-fiber/internal/entity"
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) GetAll(ctx context.Context) ([]entity.Project, error) {

	query := "SELECT * FROM projects"

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Project])

	if err != nil {
		return nil, err
	}

	return projects, nil
}
