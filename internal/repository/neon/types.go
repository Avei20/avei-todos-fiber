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
	}

	dbImpl struct {
		client *pgx.Conn
	}
)
