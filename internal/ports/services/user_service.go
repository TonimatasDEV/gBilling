package services

import (
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(email string, password string) error {
	hashedBytes, err := hashPassword(password)

	if err != nil {
		return err
	}

	user := domain.User{
		ID:             1,
		Email:          email,
		HashedPassword: string(hashedBytes),
	}

	err = s.repo.Save(user)
	return err
}

func (s *UserService) Login(email, password string) (string, bool) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", false
	}

	return "x", user.ComparePassword(password) // TODO: Generate JWT token.
}

func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}
