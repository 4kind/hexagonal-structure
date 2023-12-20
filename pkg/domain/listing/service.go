package listing

import (
	"errors"
)

// ErrNotFound is used when a user could not be found.
var ErrNotFound = errors.New("user not found")

// Repository provides access to the user and review storage.
type Repository interface {
	// GetUser returns the user with given ID.
	GetUser(string) (User, error)
	// GetAllUsers returns all users saved in storage.
	GetAllUsers() []User
}

// Service provides user and review listing operations.
type Service interface {
	GetUser(string) (User, error)
	GetUsers() []User
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetUsers returns all users
func (s *service) GetUsers() []User {
	return s.r.GetAllUsers()
}

// GetUser returns a user
func (s *service) GetUser(id string) (User, error) {
	return s.r.GetUser(id)
}
