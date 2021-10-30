package organization

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;unique;"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (org *Organization) BeforeCreate() error {
	(*org).ID = uuid.New()
	return nil
}
