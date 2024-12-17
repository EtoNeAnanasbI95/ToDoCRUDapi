package tests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"

	_ "github.com/lib/pq"
)

func SetUpTestDD() (*sqlx.DB, error) {
	dsn := os.Getenv("TEST_DB_DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test db: %v", err)
	}
	return db, nil
}
