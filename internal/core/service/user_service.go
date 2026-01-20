package service

import (
	"GoWebapitest/internal/core/domain/models"
	"GoWebapitest/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(r port.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}
func (s *UserService) AddUser(nom, email, motdepasse, role string) error {
	return s.repo.AddUser(nom, email, motdepasse, role)
}
func (s *UserService) ModifyStatus(id string, isactive bool) error {
	return s.repo.ModifyStatus(id, isactive)
}