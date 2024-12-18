package tests

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var tr *repository.TasksRepository

func TestMain(m *testing.M) {
	tableName = "users"
	query = `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		is_completed BOOLEAN DEFAULT FALSE
	);`
	code := m.Run()
	os.Exit(code)
}

func TestTasksRepositoryCreate(t *testing.T) {
	db := Setup(t, query, &tr, repository.NewTasksRepository)
	defer Teardown(t, db, tableName)
	user := &models.User{
		Name:  "testCreate",
		Email: "testCreate",
	}
	id, err := ur.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, id)

	var createdUser models.User
	err = db.Get(&createdUser, "SELECT * FROM tasks WHERE id=$1", id)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Name, createdUser.Name)
}

func TestTasksRepositoryGet(t *testing.T) {
	db := Setup(t, query, &tr, repository.NewTasksRepository)
	defer Teardown(t, db, tableName)
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

func TestTasksRepositoryGetAll(t *testing.T) {
	db := Setup(t, query, &tr, repository.NewTasksRepository)
	defer Teardown(t, db, tableName)
	setUpTasks := &[]models.User{
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
	ids := make([]int, 0, len(*setUpTasks))

	for _, user := range *setUpTasks {
		id, _ := ur.Create(&user)
		ids = append(ids, id)
	}
	users, err := ur.GetAll()
	assert.NoError(t, err)
	for index, _ := range users {
		assert.Equal(t, users[index].Id, ids[index])
	}
}

func TestTasksRepositoryUpdate(t *testing.T) {
	db := Setup(t, query, &tr, repository.NewTasksRepository)
	defer Teardown(t, db, tableName)
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

func TestTasksRepositoryDelete(t *testing.T) {
	db := Setup(t, query, &tr, repository.NewTasksRepository)
	defer Teardown(t, db, tableName)
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
