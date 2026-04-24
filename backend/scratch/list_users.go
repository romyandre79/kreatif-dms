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

	rows, err := conn.Query(ctx, "SELECT email, status, role_id FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users in DB:")
	for rows.Next() {
		var email, status string
		var roleID []byte // uuid
		rows.Scan(&email, &status, &roleID)
		fmt.Printf("- %s (Status: %s, RoleID: %x)\n", email, status, roleID)
	}
}
