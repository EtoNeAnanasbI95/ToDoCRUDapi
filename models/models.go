package models

import "errors"

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

type UserInput struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type TaskInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsCompleted *bool   `json:"is_completed"`
}

func (i *UserInput) Validate() error {
	if i.Email == nil && i.Name == nil {
		return errors.New("email or Name is required")
	}
	return nil
}

func (i *TaskInput) Validate() error {
	if i.Description == nil && i.Name == nil && i.IsCompleted == nil {
		return errors.New("description or Name or IsCompleted is required")
	}
	return nil
}
