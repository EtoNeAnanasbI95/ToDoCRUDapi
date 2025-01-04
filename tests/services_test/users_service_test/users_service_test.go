package users_service_test

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/service"
	"testing"
)

var us *service.UsersService

const tableName = "users"
const query = `
	CREATE TABLE IF NOT EXISTS users
	(
		id    SERIAL PRIMARY KEY,
		name  TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`

func TestGet(t *testing.T) {
}

func TestGetAll(t *testing.T) {

}

func TestCreate(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
