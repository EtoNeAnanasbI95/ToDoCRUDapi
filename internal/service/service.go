package service

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
)

type Users interface {
	Create(user *models.User) (int, error)
	Get(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(id int, user *models.UserInput) error
	Delete(id int) error
}

type Tasks interface {
	Create(userId int64, user *models.Task) (int, error)
	Get(userId int64, id int) (*models.Task, error)
	GetAll(userId int64) ([]models.Task, error)
	Update(userId int64, id int, task *models.TaskInput) error
	Delete(userId int64, id int) error
}

type Service struct {
	Users
	Tasks
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(r.Users),
		Tasks: NewTasksService(r.Tasks, r.UsersTasks),
	}
}
