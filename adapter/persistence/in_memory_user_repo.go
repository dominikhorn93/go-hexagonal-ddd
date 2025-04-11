package persistence

import (
	"errors"

	"hexagonal-example/application/port/out"
	"hexagonal-example/domain"
)

// InMemoryUserRepo implements the outbound port (UserRepository).
type InMemoryUserRepo struct {
	storage map[string]domain.User
}

// NewInMemoryUserRepo is a Wire provider function returning a concrete repo.
func NewInMemoryUserRepo() out.UserRepository {
	return &InMemoryUserRepo{
		storage: make(map[string]domain.User),
	}
}

func (r *InMemoryUserRepo) Save(user domain.User) error {
	if _, exists := r.storage[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.storage[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id string) (domain.User, error) {
	user, ok := r.storage[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}
