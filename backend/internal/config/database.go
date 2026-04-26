package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase(databaseURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("ERROR: Unable to parse database URL: %v\n", err)
		return nil, err
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Printf("WARNING: Unable to connect to database: %v\n", err)
		return nil, err
	}

	// Test connection
	err = pool.Ping(context.Background())
	if err != nil {
		log.Printf("WARNING: Database ping failed: %v\n", err)
		return pool, err // Return pool anyway, it might connect later
	}

	log.Println("Database connection established")
	return pool, nil
}
