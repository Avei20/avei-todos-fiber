package pgx

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context, connectionString string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, connectionString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connected to database with connection string %v\n", connectionString)
	return conn
}
