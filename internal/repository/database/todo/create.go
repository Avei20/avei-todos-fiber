package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *RepositoryImpl) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	projectId := pgtype.Text{}
	projectId.Scan(todo.ProjectId)
	projectIdValue, _ := projectId.Value()

	parentId := pgtype.Text{}
	parentId.Scan(todo.ParentId)
	parentIdValue, _ := parentId.Value()

	query := `INSERT
	INTO  todos (id, name, parent_id, project_id, is_done)
	VALUES ($1, $2, $3, $4, $5)`

	// rows, err := r.db.Query(ctx, query, todo.Id, todo.Name, parentIdValue, projectIdValue, todo.IsDone)
	rows, err := r.db.Exec(ctx, query, todo.Id, todo.Name, parentIdValue, projectIdValue, todo.IsDone)

	log.Println(rows)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
