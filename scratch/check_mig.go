package main
import (
	"fmt"
	"log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)
func main() {
	m, err := migrate.New("file://../backend/db/migrations", "postgresql://postgres:postgres@localhost:5432/kreatif_dms?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	v, d, err := m.Version()
	fmt.Printf("Version: %v, Dirty: %v, Err: %v\n", v, d, err)
}
