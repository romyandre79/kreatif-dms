package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:123456@localhost:5432/kreatif_dms?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	fmt.Println("Applying migration: Add avatar and signature fields...")
	_, err = conn.Exec(ctx, `
		ALTER TABLE users ADD COLUMN IF NOT EXISTS avatar_url TEXT;
		ALTER TABLE users ADD COLUMN IF NOT EXISTS signature_url TEXT;
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration successful!")
}
