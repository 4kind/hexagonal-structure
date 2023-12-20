package json

import (
	"encoding/json"
	"github.com/sdomino/scribble"
	"hexagonal-structure/pkg/core/storage"
	"hexagonal-structure/pkg/domain/adding"
	"hexagonal-structure/pkg/domain/deleting"
	"hexagonal-structure/pkg/domain/listing"
	"log"
	"path"
	"runtime"
	"time"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionUser identifier for the JSON collection of users
	CollectionUser = "users"
)

// Storage stores user data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddUser saves the given user to the repository
func (s *Storage) AddUser(b adding.User) error {
	id, err := storage.GetID("user")
	if err != nil {
		log.Fatal(err)
	}

	newB := User{
		ID:       id,
		Created:  time.Now(),
		FirsName: b.FirstName,
		LastName: b.LastName,
		Email:    b.Email,
		Phone:    b.Phone,
	}

	if err := s.db.Write(CollectionUser, newB.ID, newB); err != nil {
		return err
	}
	return nil
}

// GetUser Get returns a user with the specified ID
func (s *Storage) GetUser(id string) (listing.User, error) {
	var b User
	var user listing.User

	if err := s.db.Read(CollectionUser, id, &b); err != nil {
		// err handling omitted for simplicity
		return user, listing.ErrNotFound
	}

	user.ID = b.ID
	user.FirstName = b.FirsName
	user.LastName = b.LastName
	user.Email = b.Email
	user.Phone = b.Phone
	user.Created = b.Created

	return user, nil
}

// GetAllUsers GetAll returns all users
func (s *Storage) GetAllUsers() []listing.User {
	list := []listing.User{}

	records, err := s.db.ReadAll(CollectionUser)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var b User
		var user listing.User

		if err := json.Unmarshal([]byte(r), &b); err != nil {
			// err handling omitted for simplicity
			return list
		}

		user.ID = b.ID
		user.FirstName = b.FirsName
		user.LastName = b.LastName
		user.Email = b.Email
		user.Phone = b.Phone
		user.Created = b.Created

		list = append(list, user)
	}

	return list
}

// DeleteUser Delete returns a user with the specified ID
func (s *Storage) DeleteUser(id string) error {
	if err := s.db.Delete(CollectionUser, id); err != nil {
		return deleting.ErrNotFound
	}

	return nil
}
