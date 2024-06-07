package project

import (
	"avei-todos-fiber/internal/entity"
	"context"
	"log"
)

func (r *RepositoryImpl) Create(ctx context.Context, project entity.Project) (*entity.Project, error) {
	query := `INSERT
		INTO projects (id, name, description, deleted)
		VALUES ($1, $2, $3, false)`

	rows, err := r.db.Query(ctx, query, project.Id, project.Name, project.Description)

	log.Println(rows)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
