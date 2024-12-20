package tests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var tableName = "users"
var query = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	);`

func SetUpTestDD() (*sqlx.DB, error) {
	dsn := os.Getenv("TEST_DB_DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test db: %v", err)
	}
	return db, nil
}

func Teardown(t *testing.T, db *sqlx.DB, tables ...string) {
	if db != nil {
		if len(tables) > 1 {
			for table := range tables {
				_, err := db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE", table))
				if err != nil {
					t.Errorf("Failed to drop test table: %v", err)
				}
			}
		} else {
			_, err := db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE", tables[0]))
			if err != nil {
				t.Errorf("Failed to drop test table: %v", err)
			}
		}
		_ = db.Close()
	}
}

type initRepoFn[Repo any] func(db *sqlx.DB) Repo

func Setup[Repository any](t *testing.T, query string, repo **Repository, init initRepoFn[*Repository]) *sqlx.DB {
	db, err := SetUpTestDD()
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}
	*repo = init(db)
	_, err = db.Exec(query)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}
	return db
}
