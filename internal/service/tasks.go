package service

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
)

const tasksErrorPrefix = "[tasks_service]"

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
	userTasks, err := us.rut.GetWhereUid(userId)
	if err != nil {
		return nil, fmt.Errorf("%s: sometimes went wrong", tasksErrorPrefix)
	}
	for _, record := range userTasks {
		if record.TaskId == id {
			return us.rt.Get(id)
		}
	}
	return nil, fmt.Errorf("%s: task not found", tasksErrorPrefix)
}

func (us *TasksService) GetAll(userId int) ([]models.Task, error) {
	userTasks, err := us.rut.GetWhereUid(userId)
	if err != nil {
		return nil, fmt.Errorf("%s: sometimes went wrong, %w", tasksErrorPrefix, err)
	}
	tasks := make([]models.Task, 0, len(userTasks))
	for _, record := range userTasks {
		task, err := us.rt.Get(record.TaskId)
		if err != nil {
			return nil, fmt.Errorf("%s: sometimes went wrong, %w", tasksErrorPrefix, err)
		}
		tasks = append(tasks, *task)
	}
	if len(tasks) == 0 {
		return nil, fmt.Errorf("%s: tasks not found", tasksErrorPrefix)
	}
	return tasks, nil
}

func (us *TasksService) Update(userId int, id int, task *models.TaskInput) error {
	if err := task.Validate(); err != nil {
		return fmt.Errorf("%s: %w", tasksErrorPrefix, err)
	}
	userTasks, err := us.rut.GetWhereUid(userId)
	if err != nil {
		return fmt.Errorf("%s: user with id -- %d, not found, %w", tasksErrorPrefix, userId, err)
	}
	for _, userTask := range userTasks {
		if userTask.TaskId == id {
			return us.rt.Update(id, task)
		}
	}
	return fmt.Errorf("%s: task not found", tasksErrorPrefix)
}

func (us *TasksService) Delete(userId int, id int) error {
	userTasks, err := us.rut.GetWhereUid(userId)
	if err != nil {
		return fmt.Errorf("%s: user with id -- %d, not found, %w", tasksErrorPrefix, userId, err)
	}
	for _, userTask := range userTasks {
		if userTask.TaskId == id {
			return us.rt.Delete(id)
		}
	}
	return fmt.Errorf("%s: task not found", tasksErrorPrefix)
}
