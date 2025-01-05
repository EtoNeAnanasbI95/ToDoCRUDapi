package service

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
)

const tasksErrorPrefix = "[users_repository]"

type TasksService struct {
	rt  repository.Tasks
	rut repository.UsersTasks
}

func NewTasksService(rt repository.Tasks, rut repository.UsersTasks) *TasksService {
	return &TasksService{
		rt:  rt,
		rut: rut,
	}
}

func (us *TasksService) Create(userId int, task *models.Task) (int, error) {
	id, err := us.rt.Create(task)
	if err != nil {
		return 0, err
	}
	_, err = us.rut.Create(&models.UsersTasks{
		UserId: userId,
		TaskId: id,
	})
	return id, nil
}

func (us *TasksService) Get(userId int, id int) (*models.Task, error) {
	user, err := us.rt.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *TasksService) GetAll(userId int) ([]models.Task, error) {
	users, err := us.rt.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *TasksService) Update(userId int, id int, user *models.TaskInput) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	return us.rt.Update(id, user)
}

func (us *TasksService) Delete(userId int, id int) error {
	return us.rt.Delete(id)
}
