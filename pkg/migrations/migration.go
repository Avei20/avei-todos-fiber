package migrations

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	m, err := migrate.New(
		"file://pkg/migrations",
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		fmt.Println("Error on migration new")
		panic(err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		fmt.Println("Error on migration up")
		panic(err)
	}

	fmt.Println("Migration completed")
}
