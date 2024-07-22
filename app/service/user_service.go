package service

import (
	"TheBoys/app/model/request"
	"TheBoys/domain"
)

func NewUserServices(repo domain.UserRepository) domain.UserService {
	return &userService{repo}
}

type userService struct {
	repo domain.UserRepository
}

func (s *userService) CreateUser(req request.UserCreationRequest) error {

	err := s.repo.CreateUser(req)

	if err != nil {
		return err
	}
	return nil
}
