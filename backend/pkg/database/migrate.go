package database

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string, path string, command string) {
	m, err := migrate.New(
		"file://"+path,
		dbURL,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	switch command {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migration UP failed: %v", err)
		}
		log.Println("Database migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migration DOWN failed: %v", err)
		}
		log.Println("Database migrations reverted successfully")
	default:
		log.Fatalf("Unknown migration command: %s", command)
	}
}
