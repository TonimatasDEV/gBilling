package repositories

import "github.com/TonimatasDEV/BillingPanel/internal/domain"

type UserRepository interface {
	Save(user domain.User)
}
