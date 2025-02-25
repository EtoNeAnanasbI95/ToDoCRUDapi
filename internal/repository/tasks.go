package repository

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

const tasksErrorPrefix = "[tasks_repository]"

type TasksRepository struct {
	db *sqlx.DB
}

func NewTasksRepository(db *sqlx.DB) *TasksRepository {
	return &TasksRepository{
		db: db,
	}
}

func (ur *TasksRepository) Create(task *models.Task) (int, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", tasksErrorPrefix, err)
	}
	var taskId int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (name, description, is_completed) VALUES ($1, $2, $3) RETURNING id", tasksTable)
	row := tx.QueryRow(createTaskQuery, task.Name, task.Description, task.IsCompleted)
	err = row.Scan(&taskId)
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("%s: %w", tasksErrorPrefix, err)
	}
	return taskId, tx.Commit()
}

func (ur *TasksRepository) Get(id int) (*models.Task, error) {
	var task models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tasksTable)
	if err := ur.db.Get(&task, query, id); err != nil {
		return nil, fmt.Errorf("%s: %w", tasksErrorPrefix, err)
	}
	return &task, nil
}

func (ur *TasksRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	query := fmt.Sprintf("SELECT * FROM %s", tasksTable)
	if err := ur.db.Select(&tasks, query); err != nil {
		return nil, fmt.Errorf("%s: %w", tasksErrorPrefix, err)
	}
	return tasks, nil
}

func (ur *TasksRepository) Update(id int, task *models.TaskInput) error {
	setValues := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	argsId := 1
	if task.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name = $%d", argsId))
		argsId++
		args = append(args, *task.Name)
	}
	if task.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argsId))
		argsId++
		args = append(args, *task.Description)
	}
	if task.IsCompleted != nil {
		setValues = append(setValues, fmt.Sprintf("is_completed  = $%d", argsId))
		argsId++
		args = append(args, *task.IsCompleted)
	}

	setQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", tasksTable, strings.Join(setValues, ", "), argsId)
	args = append(args, id)

	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(setQuery, args...)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", tasksErrorPrefix, err)
	}
	return tx.Commit()
}

func (ur *TasksRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tasksTable)
	_, err := ur.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", tasksErrorPrefix, err)
	}
	return nil
}
