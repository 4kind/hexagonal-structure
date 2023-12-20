package memory

import (
	"hexagonal-structure/pkg/core/storage"
	"hexagonal-structure/pkg/domain/adding"
	"hexagonal-structure/pkg/domain/deleting"
	"hexagonal-structure/pkg/domain/listing"
	"log"
	"time"
)

// Storage Memory storage keeps data in memory
type Storage struct {
	users []User
}

// AddUser Add saves the given user to the repository
func (m *Storage) AddUser(b adding.User) error {
	id, err := storage.GetID("user")
	if err != nil {
		log.Fatal(err)
	}

	newB := User{
		ID:        id,
		Created:   time.Now(),
		FirstName: b.FirstName,
		LastName:  b.LastName,
		Email:     b.Email,
		Phone:     b.Phone,
	}
	m.users = append(m.users, newB)

	return nil
}

// GetUser Get returns a user with the specified ID
func (m *Storage) GetUser(id string) (listing.User, error) {
	var user listing.User

	for i := range m.users {

		if m.users[i].ID == id {
			user.ID = m.users[i].ID
			user.FirstName = m.users[i].FirstName
			user.LastName = m.users[i].LastName
			user.Email = m.users[i].Email
			user.Phone = m.users[i].Phone
			user.Created = m.users[i].Created

			return user, nil
		}
	}

	return user, listing.ErrNotFound
}

// GetAllUsers GetAll return all users
func (m *Storage) GetAllUsers() []listing.User {
	var users []listing.User

	for i := range m.users {

		user := listing.User{
			ID:        m.users[i].ID,
			FirstName: m.users[i].FirstName,
			LastName:  m.users[i].LastName,
			Email:     m.users[i].Email,
			Phone:     m.users[i].Phone,
			Created:   m.users[i].Created,
		}

		users = append(users, user)
	}

	return users
}

// DeleteUser Delete returns a user with the specified ID
func (m *Storage) DeleteUser(id string) error {

	for i := range m.users {

		if m.users[i].ID == id {
			// Remove user from the slice
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}

	return deleting.ErrNotFound
}
