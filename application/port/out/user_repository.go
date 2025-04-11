package out

import "hexagonal-example/domain"

// UserRepository is an outbound port (what the application needs from storage).
type UserRepository interface {
	Save(user domain.User) error
	FindByID(id string) (domain.User, error)
}
