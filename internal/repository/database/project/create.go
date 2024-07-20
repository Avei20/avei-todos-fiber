package project

import (
	"avei-todos-fiber/internal/entity"
	"context"
	"log"
)

func (r *RepositoryImpl) Create(ctx context.Context, project *entity.Project) (*entity.Project, error) {
	query := `INSERT
		INTO projects (id, name, description, deleted, code)
		VALUES ($1, $2, $3, false, $4)`

	log.Println(query)
	log.Printf("%+v", project)
	rows, err := r.db.Exec(ctx, query, project.Id, project.Name, project.Description, project.Code)

	log.Println(rows)
	if err != nil {
		return nil, err
	}

	return project, nil
}
