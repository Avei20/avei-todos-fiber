package todo

import (
	"avei-todos-fiber/internal/entity"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *RepositoryImpl) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	log.Printf("%+v", *todo.ProjectId)
	projectId := pgtype.Text{}
	projectId.Scan(*todo.ProjectId)
	projectIdValue, err := projectId.Value()

	if err != nil {
		log.Println(err)
	}

	parentId := pgtype.Text{}
	parentId.Scan(*todo.ParentId)
	parentIdValue, err := parentId.Value()

	if err != nil {
		log.Println(err)
	}

	log.Printf("ProjectId Value:", projectIdValue)
	log.Printf("ParentId Value:", parentIdValue)

	// query := `INSERT
	// INTO  todos (id, name, parent_id, project_id, is_done)
	// VALUES ($1, $2, $3, $4, $5)`

	// rows, err := r.db.Query(ctx, query, todo.Id, todo.Name, parentIdValue, projectIdValue, todo.IsDone)
	// _, err = r.db.Exec(ctx, query, todo.Id, todo.Name, parentIdValue, projectIdValue, todo.IsDone)

	// if err != nil {
	// 	return nil, err
	// }

	return todo, nil
}
