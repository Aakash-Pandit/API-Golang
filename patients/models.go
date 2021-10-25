package patients

import "github.com/google/uuid"

type Patient struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Contact   string    `json:"contact"`
}
