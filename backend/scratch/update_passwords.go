package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:123456@localhost:5432/kreatif_dms?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	fmt.Printf("Updating all users to password: %s\n", string(hashed))

	_, err = conn.Exec(ctx, "UPDATE users SET password_hash = $1", string(hashed))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done! All users now use 'password' as their password.")
}
