package patients

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Patient struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Contact   string    `json:"contact"`
}

type Medicine struct {
	gorm.Model
	Name string `json:"name"`
	Cost uint   `json:"cost"`
}
