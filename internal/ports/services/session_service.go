package services

import (
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
)

type SessionService struct {
	sessionRepo repositories.SessionRepository
}

func NewSessionService(sessionRepo repositories.SessionRepository) *SessionService {
	return &SessionService{sessionRepo: sessionRepo}
}

func (s *SessionService) Remove(token string) error {
	return s.sessionRepo.Remove(token)
}

func (s *SessionService) Create(id int) (*domain.Session, error) {
	return s.sessionRepo.Create(id)
}

func (s *SessionService) Validate(token string) (*domain.Session, error) {
	return s.sessionRepo.Validate(token)
}
