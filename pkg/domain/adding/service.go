package adding

import (
	"errors"
	"hexagonal-structure/pkg/domain/listing"
)

// ErrDuplicate is used when a user already exists.
var ErrDuplicate = errors.New("user already exists")

// Service provides user adding operations.
type Service interface {
	AddUser(...User) error
	AddSampleUsers([]User)
}

// Repository provides access to user repository.
type Repository interface {
	// AddBUser saves a given user to the repository.
	AddUser(User) error
	// GetAllUsers returns all users saved in storage.
	GetAllUsers() []listing.User
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser persists the given user(s) to storage
func (s *service) AddUser(b ...User) error {
	// make sure we don't add any duplicates
	//existingUsers := s.r.GetAllUsers()
	//for _, bb := range b {
	//for _, e := range existingUsers {
	//	if bb.Abv == e.Abv &&
	//		bb.Brewery == e.Brewery &&
	//		bb.Name == e.Name {
	//		return ErrDuplicate
	//	}
	//}
	//}

	// any other validation can be done here

	for _, user := range b {
		_ = s.r.AddUser(user) // error handling omitted for simplicity
	}

	return nil
}

// AddSampleUsers adds some sample users to the database
func (s *service) AddSampleUsers(b []User) {

	// any validation can be done here

	for _, bb := range b {
		_ = s.r.AddUser(bb) // error handling omitted for simplicity
	}
}
