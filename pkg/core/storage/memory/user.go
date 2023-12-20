package memory

import (
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     int
	Created   time.Time
}
