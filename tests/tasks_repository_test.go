package tests

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ur *repository.UsersRepository

func teardown(t *testing.T, db *sqlx.DB) {
	if db != nil {
		_, err := db.Exec("DROP TABLE users CASCADE")
		if err != nil {
			t.Errorf("Failed to drop test table: %v", err)
		}
		db.Close()
	}
}

func setup(t *testing.T) *sqlx.DB {
	db, err := SetUpTestDD()
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}
	ur = repository.NewUsersRepository(db)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL
		);
	`)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}
	return db
}

func TestUsersRepositoryCreate(t *testing.T) {
	//db := setup(t)
	//defer teardown(t, db)
	const tableName = "users"
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	);`
	var ur *repository.UsersRepository
	db := Setup(t, query, &ur, repository.NewUsersRepository)
	defer Teardown(t, db, tableName)
	user := &models.User{
		Name:  "testCreate",
		Email: "testCreate",
	}
	id, err := ur.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, id)

	var createdUser models.User
	err = db.Get(&createdUser, "SELECT * FROM users WHERE id=$1", id)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Name, createdUser.Name)
}

func TestUsersRepositoryGet(t *testing.T) {
	db := setup(t)
	defer teardown(t, db)
	setUpUser := &models.User{
		Name:  "testGet",
		Email: "testGet",
	}
	id, _ := ur.Create(setUpUser)
	user, err := ur.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, setUpUser.Name, user.Name)
	assert.Equal(t, setUpUser.Email, user.Email)
}

func TestUsersRepositoryGetAll(t *testing.T) {
	db := setup(t)
	defer teardown(t, db)
	setUpUsers := &[]models.User{
		{
			Name:  "testGetAll",
			Email: "testGetAll",
		},
		{
			Name:  "testGetAll",
			Email: "testGetAll",
		},
		{
			Name:  "testGetAll",
			Email: "testGetAll",
		},
	}
	ids := make([]int, 0, len(*setUpUsers))

	for _, user := range *setUpUsers {
		id, _ := ur.Create(&user)
		ids = append(ids, id)
	}
	users, err := ur.GetAll()
	assert.NoError(t, err)
	for index, _ := range users {
		assert.Equal(t, users[index].Id, ids[index])
	}
}

func TestUsersRepositoryUpdate(t *testing.T) {
	db := setup(t)
	defer teardown(t, db)
	user := &models.User{
		Name:  "testUpdate",
		Email: "testUpdate",
	}
	id, _ := ur.Create(user)

	const updated = "_updated"
	userInput := models.User{
		Name:  updated,
		Email: updated,
	}
	err := ur.Update(id, &userInput)
	assert.NoError(t, err)
	updatedUser, _ := ur.Get(id)
	assert.Equal(t, userInput.Name, updatedUser.Name)
	assert.Equal(t, userInput.Email, updatedUser.Email)
}

func TestUsersRepositoryDelete(t *testing.T) {
	db := setup(t)
	defer teardown(t, db)
	user := &models.User{
		Name:  "test",
		Email: "test",
	}
	id, _ := ur.Create(user)

	err := ur.Delete(id)
	assert.NoError(t, err)
	_, err = ur.Get(id)
	assert.Error(t, err)
}
