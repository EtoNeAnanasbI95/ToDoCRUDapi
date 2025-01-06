package service

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
)

const usersErrorPrefix = "[users_service]"

type UsersService struct {
	r repository.Users
}

func NewUsersService(r repository.Users) *UsersService {
	return &UsersService{
		r: r,
	}
}

func (us *UsersService) Create(user *models.User) (int, error) {
	id, err := us.r.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (us *UsersService) Get(id int) (*models.User, error) {
	user, err := us.r.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UsersService) GetAll() ([]models.User, error) {
	users, err := us.r.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UsersService) Update(id int, user *models.UserInput) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("%s: %w", usersErrorPrefix, err)
	}
	return us.r.Update(id, user)
}

func (us *UsersService) Delete(id int) error {
	return us.r.Delete(id)
}
