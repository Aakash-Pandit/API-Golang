package patients

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;unique;"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"unique;"`
	Contact   string    `json:"contact"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
}

func (patient *Patient) BeforeCreate() error {
	(*patient).ID = uuid.New()
	return nil
}

type Medicine struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;unique;"`
	Name     string    `json:"name"`
	Cost     uint      `json:"cost"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (medicine *Medicine) BeforeCreate() error {
	(*medicine).ID = uuid.New()
	return nil
}
