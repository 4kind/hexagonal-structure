package json

import "time"

type User struct {
	ID       string    `json:"id"`
	FirsName string    `json:"first_name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	Phone    int       `json:"phone"`
	Created  time.Time `json:"created"`
}
