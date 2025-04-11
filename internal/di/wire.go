//go:build wireinject
// +build wireinject

package di

import (
	"net/http"

	"github.com/google/wire"
	"hexagonal-example/adapter/persistence"
	"hexagonal-example/adapter/web"
	"hexagonal-example/application/service"
)

// We define provider sets for each layer/adapter
var (
	PersistenceSet = wire.NewSet(
		// Provide an in-memory repo that implements UserRepository
		persistence.NewInMemoryUserRepo,
	)

	ServiceSet = wire.NewSet(
		// Provide a user service (implements UserUseCase)
		service.NewUserService,
	)

	WebSet = wire.NewSet(
		// Provide the HTTP controller
		web.NewUserController,
	)
)

// InitializeApp is our main Wire injector. It returns something that can run.
func InitializeApp() (*http.Server, error) {
	// Combine all sets and specify how to build the final object
	wire.Build(
		PersistenceSet,
		ServiceSet,
		WebSet,
		NewHTTPServer,
	)
	return &http.Server{}, nil
}

// NewHTTPServer is a custom constructor that sets up the routes.
func NewHTTPServer(userController *web.UserController) *http.Server {
	// Register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/users", userController.CreateUserHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return srv
}
