package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Contact   string    `json:"contact"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
}

func (user *User) BeforeCreate() error {
	(*user).ID = uuid.New()
	return nil
}
