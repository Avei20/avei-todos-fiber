package neon

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type (
	DB interface {
		Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
		QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
		Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
		// Select[T any] (ctx context.Context, dest interface{}, query string, args ...interface{}) T, error
	}

	dbImpl struct {
		client *pgx.Conn
	}
)

// type Todos interface {
// 	Integer | Float | ~string
// }
