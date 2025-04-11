package domain

import (
	"errors"
	"github.com/google/uuid"
)

// User is a core domain entity.
type User struct {
	ID   string
	Name string
	Age  int
}

func NewUser(name string, age int) (User, error) {
	if name == "" {
		return User{}, errors.New("name cannot be empty")
	}
	if age < 0 {
		return User{}, errors.New("age cannot be negative")
	}

	return User{
		ID:   uuid.NewString(),
		Name: name,
		Age:  age,
	}, nil
}
