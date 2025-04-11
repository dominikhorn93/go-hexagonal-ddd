package web

import (
	"encoding/json"
	"log"
	"net/http"

	"hexagonal-example/application/port/in"
)

// UserController is an inbound adapter (HTTP) calling the inbound port (UserUseCase).
type UserController struct {
	userUseCase in.UserUseCase
}

// NewUserController is a Wire provider function.
func NewUserController(u in.UserUseCase) *UserController {
	return &UserController{userUseCase: u}
}

// CreateUserHandler handles POST /users
func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var command in.CreateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Println(command)

	if err := c.userUseCase.CreateUser(command); err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"user created"}`))
}
