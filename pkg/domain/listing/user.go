package listing

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     int       `json:"phone"`
	Created   time.Time `json:"created"`
}
