package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func MustInitDB(cs string) *sqlx.DB {
	dsn := os.Getenv("TEST_DB_DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
