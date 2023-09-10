package service

import "github.com/formulationapp/formulationapp/internal/repository"

type UserService interface {
}

type userService struct {
	userRepository repository.UserRepository
}

func newUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}
