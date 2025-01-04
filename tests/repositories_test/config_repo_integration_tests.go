package repositories_test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func SetUpTestDB() *sqlx.DB {
	dsn := os.Getenv("TEST_DB_DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Teardown(db *sqlx.DB, tables ...string) error {
	if db != nil {
		if len(tables) > 1 {
			for table := range tables {
				_, err := db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE", table))
				if err != nil {
					return err
				}
			}
		} else {
			_, err := db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE", tables[0]))
			if err != nil {
				return err
			}
		}
		_ = db.Close()
	}
	return nil
}

type initRepoFn[Repo any] func(db *sqlx.DB) Repo

func Setup[Repository any](t *testing.T, query string, repo **Repository, init initRepoFn[*Repository]) *sqlx.DB {
	db := SetUpTestDB()
	*repo = init(db)
	_, err := db.Exec(query)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}
	return db
}
