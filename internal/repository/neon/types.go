package neon

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type (
	DB interface {
		Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
		QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	}

	dbImpl struct {
		client *pgx.Conn
	}
)
