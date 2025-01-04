package users_repository_test

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/tests/repositories_test"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ur *repository.UsersRepository

const tableName = "users"
const query = `
	CREATE TABLE IF NOT EXISTS users
	(
		id    SERIAL PRIMARY KEY,
		name  TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`

func TestCreate(t *testing.T) {
	db := repositories_test.Setup(t, query, &ur, repository.NewUsersRepository)
	defer func() {
		assert.Nil(t, repositories_test.Teardown(db, tableName))
	}()
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

func TestGet(t *testing.T) {
	db := repositories_test.Setup(t, query, &ur, repository.NewUsersRepository)
	defer func() {
		assert.Nil(t, repositories_test.Teardown(db, tableName))
	}()
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

func TestGetAll(t *testing.T) {
	db := repositories_test.Setup(t, query, &ur, repository.NewUsersRepository)
	defer func() {
		assert.Nil(t, repositories_test.Teardown(db, tableName))
	}()
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
	for index := range users {
		assert.Equal(t, users[index].Id, ids[index])
	}
}

func TestUpdate(t *testing.T) {
	db := repositories_test.Setup(t, query, &ur, repository.NewUsersRepository)
	defer func() {
		assert.Nil(t, repositories_test.Teardown(db, tableName))
	}()
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

func TestDelete(t *testing.T) {
	db := repositories_test.Setup(t, query, &ur, repository.NewUsersRepository)
	defer func() {
		assert.Nil(t, repositories_test.Teardown(db, tableName))
	}()
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
