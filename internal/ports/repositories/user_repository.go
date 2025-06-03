package repositories

import "github.com/TonimatasDEV/BillingPanel/internal/domain"

type UserRepository interface {
	Create(user domain.User) error
	GetByEmail(email string) (domain.User, error)
}
