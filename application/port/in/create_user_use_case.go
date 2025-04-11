package in

// UserUseCase is an inbound port: the interface for user-related use cases.
type UserUseCase interface {
	CreateUser(cmd CreateUserCommand) error
}

type CreateUserCommand struct {
	Name string
	Age  int
}
