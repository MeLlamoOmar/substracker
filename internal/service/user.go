package service

import (
	"SubsTracker/internal/models"
	"SubsTracker/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
