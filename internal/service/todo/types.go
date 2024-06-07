package todo

import "avei-todos-fiber/internal/repository/database/todo"

type (
	Service     interface{}
	ServiceImpl struct {
		repo todo.Repository
	}
)
