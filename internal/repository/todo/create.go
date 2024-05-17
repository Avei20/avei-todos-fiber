package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"
	"log"
)

func (r *RepositoryImpl) Create(ctx context.Context, todo entity.Todo) (*entity.Todo, error) {
	query := `INSERT
	INTO  todos (id, name, parent_id, is_done, created_at, finish_at)
	VALUES ($1, $2, $3, $4, $5, $6)`

	rows, err := r.db.Query(ctx, query, todo.Id, todo.Name, todo.ParentId, todo.IsDone, todo.CreatedAt, todo.FinishAt)

	log.Println(rows)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
