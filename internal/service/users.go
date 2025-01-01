package service

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
)

type UsersService struct {
	r repository.Users
}

func NewUsersService(r repository.Users) *UsersService {
	return &UsersService{
		r: r,
	}
}

// TODO: завершить имплентацию сервиса
func (us *UsersService) Create(user *models.User) (int, error) {
	panic("implement me")
	return 0, nil
}

// TODO: завершить имплентацию сервиса
func (us *UsersService) Get(id int) (*models.User, error) {
	panic("implement me")
	return nil, nil
}

// TODO: завершить имплентацию сервиса
func (us *UsersService) GetAll() ([]models.User, error) {
	panic("implement me")
	return nil, nil
}

func (us *UsersService) Update(id int, user *models.User) error {
	//if err := user.Validate(); err != nil {
	//	return err
	//}
	return us.r.Update(id, user)
}

// TODO: завершить имплентацию сервиса
func (ur *UsersService) Delete(id int) error {
	panic("implement me")
	return nil
}
