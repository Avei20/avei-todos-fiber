package neon

import (
	db "avei-todos-fiber/pkg/pgx"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func NewDB(ctx context.Context, connectionString string) DB {
	db := db.ConnectDB(ctx, connectionString)

	return &dbImpl{
		client: db,
	}
}

func (d *dbImpl) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.client.Query(ctx, query, args...)
}

func (d *dbImpl) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.client.QueryRow(ctx, query, args...)
}

func (d *dbImpl) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.client.Exec(ctx, query, args...)
}

// func (d *dbImpl) Select[T] (ctx context.Context, dest interface{}, query string, args ...interface{}) T, error {
// 	data, err := d.client.Query(ctx, query, args...)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer data.Close()

// 	parsed, err := pgx.CollectRows(data, pgx.RowToStructByName([T]))
// 	return parsed, err
// }
