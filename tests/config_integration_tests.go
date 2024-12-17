package tests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func SetUpTestDD() *sqlx.DB {
	dsn := os.Getenv("TEST_DB_DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to test db: %v", err))
	}
	return db
}
