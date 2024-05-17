package todo

import "avei-todos-fiber/internal/repository/todo"

type (
	Service     interface{}
	ServiceImpl struct {
		repo todo.Repository
	}
)
