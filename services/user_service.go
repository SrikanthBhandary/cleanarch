package services

import (
	"errors"
	"sync"

	"github.com/srikanthbhandary/cleanarch/models"
)

var once sync.Once

type UserService interface {
	Create(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	Validate(user *models.User) error
}

type userService struct {
	userRepository models.UserRepo
}

var instance *userService

//NewUserService: construction function, injected by user repository
func NewUserService(r models.UserRepo) UserService {
	once.Do(func() {
		instance = &userService{
			userRepository: r,
		}
	})
	return instance
}

func (*userService) Validate(user *models.User) error {
	if user == nil {
		err := errors.New("The user is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("The name of user is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("The email of user is empty")
		return err
	}
	return nil
}

func (u *userService) Create(user *models.User) (*models.User, error) {
	return u.userRepository.Save(user)
}
func (u *userService) FindAll() ([]models.User, error) {
	return u.userRepository.FindAll()
}
