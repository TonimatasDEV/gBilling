package services

import (
	"errors"
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService struct {
	userRepo       repositories.UserRepository
	sessionService *SessionService
}

func NewUserService(userRepo repositories.UserRepository, sessionService *SessionService) *UserService {
	return &UserService{userRepo: userRepo, sessionService: sessionService}
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

	err = s.userRepo.Create(user)
	return err
}

func (s *UserService) Auth(token string) (*domain.Session, error) {
	validate, err := s.sessionService.Validate(token)
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	return validate, nil
}

func (s *UserService) Login(rawUser domain.RawUser) (string, error) {
	user, err := s.userRepo.GetByEmail(rawUser.Email)
	if err != nil {
		return "", err
	}

	err = user.ComparePassword(rawUser.Password)
	if err != nil {
		return "", err
	}

	session, err := s.sessionService.Create(user.ID)

	if err != nil {
		return "", err
	}

	return session.Token, nil
}

func (s *UserService) Logout(token string) error {
	return s.sessionService.Remove(token)
}

func (s *UserService) GetRawUser(r *http.Request) (domain.RawUser, error) {
	return domain.CreateRawUser(r)
}

func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}
