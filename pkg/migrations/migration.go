package migrations

import (
	"log"
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
		log.Println("Error on migration new")
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Println("Error on migration up")
		log.Fatal(err)
	}
}
