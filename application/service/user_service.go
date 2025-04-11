package service

import (
	"hexagonal-example/application/port/in"
	"hexagonal-example/application/port/out"
	"hexagonal-example/domain"
)

// userService implements the inbound port (UserUseCase).
type userService struct {
	userRepo out.UserRepository
}

// NewUserService is a Wire provider function. It returns the inbound port
// and requires an outbound port (UserRepository).
func NewUserService(repo out.UserRepository) in.UserUseCase {
	return &userService{userRepo: repo}
}

func (s *userService) CreateUser(cmd in.CreateUserCommand) error {
	user, err := domain.NewUser(cmd.Name, cmd.Age)
	if err != nil {
		return err
	}

	return s.userRepo.Save(user)
}
