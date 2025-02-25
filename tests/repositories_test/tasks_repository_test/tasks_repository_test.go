package tasks_repository_test

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/tests/repositories_test"
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

func TestCreate(t *testing.T) {
	db := repositories_test.Setup(t, query, &tr, repository.NewTasksRepository)
	defer func() {
		assert.NoError(t, repositories_test.Teardown(db, tableName))
	}()
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

func TestGet(t *testing.T) {
	db := repositories_test.Setup(t, query, &tr, repository.NewTasksRepository)
	defer func() {
		assert.NoError(t, repositories_test.Teardown(db, tableName))
	}()

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

func TestGetAll(t *testing.T) {
	db := repositories_test.Setup(t, query, &tr, repository.NewTasksRepository)
	defer func() {
		assert.NoError(t, repositories_test.Teardown(db, tableName))
	}()

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

func TestUpdate(t *testing.T) {
	db := repositories_test.Setup(t, query, &tr, repository.NewTasksRepository)
	defer func() {
		assert.NoError(t, repositories_test.Teardown(db, tableName))
	}()

	model := &models.Task{
		Name:        "testCreate",
		Description: "testCreate",
	}
	id, _ := tr.Create(model)
	testDataName := "testName"
	testDataDescription := "testDescription"
	testDataIsCompleted := false
	modelInput := &models.TaskInput{
		Name:        &testDataName,
		Description: &testDataDescription,
		IsCompleted: &testDataIsCompleted,
	}
	err := tr.Update(id, modelInput)
	assert.NoError(t, err)
	updatedUser, err := tr.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, *modelInput.Name, updatedUser.Name)

	testDataName = "testName2"
	testDataDescription = "testDescription2"
	modelInput = &models.TaskInput{
		Name:        &testDataName,
		Description: &testDataDescription,
	}
	err = tr.Update(id, modelInput)
	assert.NoError(t, err)
	updatedUser, err = tr.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, *modelInput.Name, updatedUser.Name)
}

func TestDelete(t *testing.T) {
	db := repositories_test.Setup(t, query, &tr, repository.NewTasksRepository)
	defer func() {
		assert.NoError(t, repositories_test.Teardown(db, tableName))
	}()

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
