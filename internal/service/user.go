package service

import "github.com/formulationapp/formulationapp/internal/repository"

type UserService interface {
	Login(email string, password string) (string, error)
	Register(email string, password string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) Login(email string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) Register(email string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func newUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}
