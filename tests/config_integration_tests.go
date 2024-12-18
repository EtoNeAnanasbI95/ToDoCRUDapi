package tests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"

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

func Teardown(t *testing.T, db *sqlx.DB, table string) {
	if db != nil {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE", table))
		if err != nil {
			t.Errorf("Failed to drop test table: %v", err)
		}
		db.Close()
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
