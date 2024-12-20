package repository

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	tasksTable      = "tasks"
	usersTasksTable = "users_tasks"
)

type Users interface {
	Create(user *models.User) (int, error)
	Get(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(id int, user *models.User) error
	Delete(id int) error
}

type Tasks interface {
	Create(user *models.Task) (int, error)
	Get(id int) (*models.Task, error)
	GetAll() ([]models.Task, error)
	Update(id int, task *models.TaskInput) error
	Delete(id int) error
}

type UsersTasks interface {
	Create(user *models.UsersTasks) (int, error)
	Get(id int) (*models.UsersTasks, error)
	GetAll() ([]models.UsersTasks, error)
	Update(id int, usersTasks *models.UsersTasks) error
	Delete(id int) error
}

type Repository struct {
	Users
	Tasks
	UsersTasks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users:      NewUsersRepository(db),
		Tasks:      NewTasksRepository(db),
		UsersTasks: NewUsersTasksRepository(db),
	}
}
