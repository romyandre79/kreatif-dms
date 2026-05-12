package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kreatif/dms-backend/internal/config"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dbPool, err := config.InitDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	rows, err := dbPool.Query(context.Background(), "SELECT id, name FROM system_modules ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("System Modules List:")
	found := false
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("- %s: %s\n", id, name)
		if id == "doc_category" {
			found = true
		}
	}

	if found {
		fmt.Println("\nSUCCESS: 'doc_category' found in database.")
	} else {
		fmt.Println("\nERROR: 'doc_category' NOT found in database.")
	}
}
