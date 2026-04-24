package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:123456@127.0.0.1:5432/kreatif_dms?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	var email, fullName, status string
	err = conn.QueryRow(context.Background(), "SELECT email, full_name, status FROM users WHERE email='admin@kreatif.id'").Scan(&email, &fullName, &status)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	fmt.Printf("User Found: %s (%s) - Status: %s\n", email, fullName, status)
}
