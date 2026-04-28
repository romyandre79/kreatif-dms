package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	dbURL := "postgresql://postgres:123456@localhost:5432/kreatif_dms?sslmode=disable"
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, "SELECT name, ldap_group FROM roles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Roles in DB:")
	for rows.Next() {
		var name string
		var ldapGroup *string
		if err := rows.Scan(&name, &ldapGroup); err != nil {
			log.Fatal(err)
		}
		lg := "NULL"
		if ldapGroup != nil {
			lg = *ldapGroup
		}
		fmt.Printf("- Name: %s, LDAP Group: %s\n", name, lg)
	}
}
