package deleting

import (
	"errors"
	"hexagonal-structure/pkg/domain/listing"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	GetUser(string) (listing.User, error)
	DeleteUser(string) error
}

type Service interface {
	DeleteUser(string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteUser(id string) error {
	return s.r.DeleteUser(id)
}
