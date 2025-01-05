package services_test

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/tests/repositories_test"
	"github.com/jmoiron/sqlx"
	"testing"
)

type initRepoFn[Repo any] func(db *sqlx.DB) Repo

func Setup[Repository any](t *testing.T, query string, repo **Repository, init initRepoFn[*Repository]) *sqlx.DB {
	db := repositories_test.SetUpTestDB()
	*repo = init(db)
	_, err := db.Exec(query)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}
	return db
}
