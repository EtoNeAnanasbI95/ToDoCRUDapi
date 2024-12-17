package repository

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UsersTasksRepository struct {
	db *sqlx.DB
}

func NewUsersTasksRepository(db *sqlx.DB) *UsersTasksRepository {
	return &UsersTasksRepository{
		db: db,
	}
}

func (ur *UsersTasksRepository) Create(user *models.UsersTasks) (int, error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return 0, err
	}
	var usersTasksId int
	createUserQuery := fmt.Sprintf("INSERT INTO %s (user_id, task_id) VALUES ($1, $2) RETURNING id)", usersTasksTable)
	row := tx.QueryRow(createUserQuery, user.UserId, user.TaskId)
	err = row.Scan(&usersTasksId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return usersTasksId, tx.Commit()
}

func (ur *UsersTasksRepository) Get(id int) (models.UsersTasks, error) {
	var usersTask models.UsersTasks
	query := fmt.Sprintf("SELECT * FROM %s WHERE id == $1", usersTasksTable)
	if err := ur.db.Select(&usersTask, query, id); err != nil {
		return usersTask, err
	}
	return usersTask, nil
}

func (ur *UsersTasksRepository) GetAll() ([]models.UsersTasks, error) {
	var usersTasks []models.UsersTasks
	query := fmt.Sprintf("SELECT * FROM %s", usersTasksTable)
	if err := ur.db.Select(&usersTasks, query); err != nil {
		return nil, err
	}
	return usersTasks, nil
}

func (ur *UsersTasksRepository) Update(id int, usersTasks *models.UsersTasks) error {
	setValues := make([]string, 2)
	args := make([]interface{}, 2)
	argsId := 1
	if usersTasks.UserId > 0 {
		setValues = append(setValues, fmt.Sprintf("user_id = $%d", argsId))
		argsId++
		args = append(args, usersTasks.UserId)
	}
	if usersTasks.TaskId > 0 {
		setValues = append(setValues, fmt.Sprintf("task_id = $%d", argsId))
		argsId++
		args = append(args, usersTasks.TaskId)
	}
	setQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", usersTasksTable, strings.Join(setValues, ", "), argsId)
	args = append(args, id)

	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(setQuery, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (ur *UsersTasksRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id == $1", usersTasksTable)
	_, err := ur.db.Exec(query, id)
	return err
}
