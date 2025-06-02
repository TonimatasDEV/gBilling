package services

import (
	"errors"
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserService struct {
	userRepo    repositories.UserRepository
	sessionRepo repositories.SessionRepository
}

func NewUserService(userRepo repositories.UserRepository, sessionRepo repositories.SessionRepository) *UserService {
	return &UserService{userRepo: userRepo, sessionRepo: sessionRepo}
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

	err = s.userRepo.Save(user)
	return err
}

func (s *UserService) Auth(r *http.Request) (*domain.Session, error) {
	cookies := r.CookiesNamed("session")

	if len(cookies) == 0 {
		return nil, errors.New("no cookies")
	}

	cookie := cookies[0]

	validate, err := s.sessionRepo.Validate(cookie.Value)
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

	session, err := s.sessionRepo.Create(user.ID)

	if err != nil {
		return "", err
	}

	return session.Token, nil
}

func (s *UserService) RemoveSession(token string) error {
	return s.sessionRepo.Remove(token)
}

func (s *UserService) GetRawUser(r *http.Request) (domain.RawUser, error) {
	return domain.CreateRawUser(r)
}

func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}
