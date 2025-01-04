package models

import "fmt"

type User struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name" binding:"required"`
	Email string `db:"email" json:"email" binding:"required"`
}

type Task struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name" binding:"required"`
	Description string `db:"description" json:"description" binding:"required"`
	IsCompleted bool   `db:"is_completed" json:"is_completed"`
}

type UsersTasks struct {
	Id     int `db:"id" json:"id"`
	UserId int `db:"user_id" json:"user_id"`
	TaskId int `db:"task_id" json:"task_id"`
}

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
