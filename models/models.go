package models

import "fmt"

// User представляет модель пользователя
type User struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name" binding:"required"`
	Email string `db:"email" json:"email" binding:"required"`
}

// Task представляет модель задачи
type Task struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
	IsCompleted bool   `db:"is_completed" json:"is_completed"`
}

// UsersTasks представляет модель задачи пользователя, как обращение к таблице MtM
type UsersTasks struct {
	Id     int   `db:"id" json:"id"`
	UserId int64 `db:"user_id" json:"user_id"`
	TaskId int   `db:"task_id" json:"task_id"`
}

// TaskInput представляет модель обновления задачи
type TaskInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsCompleted *bool   `json:"is_completed"`
}

func (u *TaskInput) Validate() error {
	if u.Name == nil && u.Description == nil && u.IsCompleted == nil {
		return fmt.Errorf("task lables is empty")
	}
	return nil
}

// UserInput представляет модель обновления пользователя
type UserInput struct {
	Id    *int    `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

func (u *UserInput) Validate() error {
	if u.Name == nil && u.Email == nil {
		return fmt.Errorf("user lables is empty")
	}
	return nil
}
