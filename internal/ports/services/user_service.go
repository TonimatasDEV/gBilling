package services

import (
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(email string, password string) (error, domain.User) {
	user := domain.User{
		ID:             1,
		Email:          email,
		HashedPassword: password,
	}

	err := s.repo.Save(user)
	return err, user
}
