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
	case "force":
		// command is "force <version>", but command here is just the string from Arg[2]
		// We need to handle this in main.go to pass the version as well
		log.Println("Use 'migrate force <version>' to fix dirty database")
	default:
		log.Fatalf("Unknown migration command: %s", command)
	}
}

func RunMigrationsWithVersion(dbURL string, path string, command string, version int) {
	m, err := migrate.New(
		"file://"+path,
		dbURL,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	switch command {
	case "force":
		if err := m.Force(version); err != nil {
			log.Fatalf("Migration FORCE failed: %v", err)
		}
		log.Printf("Database version forced to %d\n", version)
	case "goto":
		if err := m.Migrate(uint(version)); err != nil {
			log.Fatalf("Migration GOTO failed: %v", err)
		}
		log.Printf("Database migrated to version %d\n", version)
	}
}
