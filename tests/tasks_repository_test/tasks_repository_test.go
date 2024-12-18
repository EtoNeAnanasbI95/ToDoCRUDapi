package tasks_repository_test

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tr *repository.TasksRepository

const tableName = "tasks"
const query = `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		is_completed BOOLEAN DEFAULT FALSE
	);`

func TestTasksRepositoryCreate(t *testing.T) {
	db := tests.Setup(t, query, &tr, repository.NewTasksRepository)
	defer tests.Teardown(t, db, tableName)
	task := &models.Task{
		Name:        "testCreate",
		Description: "testCreate",
	}
	id, err := tr.Create(task)
	assert.NoError(t, err)
	assert.NotZero(t, id)

	var createdTask models.Task
	err = db.Get(&createdTask, "SELECT * FROM tasks WHERE id=$1", id)
	assert.NoError(t, err)
	assert.Equal(t, task.Name, createdTask.Name)
	assert.Equal(t, task.Description, createdTask.Description)
}

func TestTasksRepositoryGet(t *testing.T) {
	db := tests.Setup(t, query, &tr, repository.NewTasksRepository)
	defer tests.Teardown(t, db, tableName)
	setUpModel := &models.Task{
		Name:        "testCreate",
		Description: "testCreate",
	}
	id, _ := tr.Create(setUpModel)
	get, err := tr.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, setUpModel)
	assert.Equal(t, setUpModel.Name, get.Name)
}

func TestTasksRepositoryGetAll(t *testing.T) {
	db := tests.Setup(t, query, &tr, repository.NewTasksRepository)
	defer tests.Teardown(t, db, tableName)
	setUpModels := &[]models.Task{
		{
			Name:        "testGetAll",
			Description: "testGetAll",
		},
		{
			Name:        "testGetAll",
			Description: "testGetAll",
		},
		{
			Name:        "testGetAll",
			Description: "testGetAll",
		},
	}
	ids := make([]int, 0, len(*setUpModels))

	for _, model := range *setUpModels {
		id, _ := tr.Create(&model)
		ids = append(ids, id)
	}
	models, err := tr.GetAll()
	assert.NoError(t, err)
	for index := range models {
		assert.Equal(t, models[index].Id, ids[index])
	}
}

func TestTasksRepositoryUpdate(t *testing.T) {
	db := tests.Setup(t, query, &tr, repository.NewTasksRepository)
	defer tests.Teardown(t, db, tableName)
	model := &models.Task{
		Name:        "testCreate",
		Description: "testCreate",
	}
	id, _ := tr.Create(model)

	const updated = "_updated"
	modelInput := models.Task{
		Name:        updated,
		Description: updated,
	}
	err := tr.Update(id, &modelInput)
	assert.NoError(t, err)
	updatedUser, err := tr.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, modelInput.Name, updatedUser.Name)
}

func TestTasksRepositoryDelete(t *testing.T) {
	db := tests.Setup(t, query, &tr, repository.NewTasksRepository)
	defer tests.Teardown(t, db, tableName)
	model := &models.Task{
		Name:        "test",
		Description: "test",
	}
	id, _ := tr.Create(model)

	err := tr.Delete(id)
	assert.NoError(t, err)
	_, err = tr.Get(id)
	assert.Error(t, err)
}
