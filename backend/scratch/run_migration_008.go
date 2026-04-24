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

	fmt.Println("Applying migration: Add manager_doc_controller role and user...")
	_, err = conn.Exec(ctx, `
		-- Add manager_doc_controller role
		INSERT INTO roles (name, description) VALUES 
		('manager_doc_controller', 'Head of Document Control with full operational management access')
		ON CONFLICT (name) DO NOTHING;

		-- Seed a user for manager_doc_controller
		INSERT INTO users (email, password_hash, full_name, role_id, status)
		SELECT 'manager_dc@kreatif.id', '$2a$10$qXhx1.rvn8kO.Cg6WqO2aO/lwyVsVQkhuaVhdW2kjPpQp/hITuzY.', 'Head of Document Control', id, 'approved'
		FROM roles WHERE name = 'manager_doc_controller'
		ON CONFLICT (email) DO NOTHING;
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration successful!")
}
