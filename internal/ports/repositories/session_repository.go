package repositories

import (
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
)

type SessionRepository interface {
	Create(userId int) (*domain.Session, error)
	Remove(token string) error
	Validate(r string) (*domain.Session, error)
	GetById(id int64) (*domain.Session, error)
}
