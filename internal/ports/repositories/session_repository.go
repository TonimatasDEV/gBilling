package repositories

import (
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
)

type SessionRepository interface {
	Create(userID int) (*domain.Session, error)
	Remove(token string) error
	Validate(r string) (*domain.Session, error)
	GetByID(id int64) (*domain.Session, error)
}
