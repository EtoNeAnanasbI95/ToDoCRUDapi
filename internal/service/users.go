package service

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
)

type UsersService struct {
	r *repository.Repository
}

func NewUsersService(r *repository.Repository) *UsersService {
	return &UsersService{
		r: r,
	}
}

// TODO: завершить имплентацию сервиса
func (ur *UsersService) Create(user *models.User) (int, error) {
	panic("implement me")
	return 0, nil
}

// TODO: завершить имплентацию сервиса
func (ur *UsersService) Get(id int) (string, error) {
	panic("implement me")
	return "", nil
}

// TODO: завершить имплентацию сервиса
func (ur *UsersService) GetAll() ([]string, error) {
	panic("implement me")
	return nil, nil
}

func (ur *UsersService) Update(id int, user *models.UserInput) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return ur.r.Users.Update(id, user)
}

// TODO: завершить имплентацию сервиса
func (ur *UsersService) Delete(id int) error {
	panic("implement me")
	return nil
}
