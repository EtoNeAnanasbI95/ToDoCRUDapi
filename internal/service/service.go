package service

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/models"
)

type Users interface {
	Create(user *models.User) (int, error)
	Get(int) (string, error)
	GetAll() ([]string, error)
	Update(int) error
	Delete(int) error
}

type Service struct {
	Users
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(r),
	}
}
