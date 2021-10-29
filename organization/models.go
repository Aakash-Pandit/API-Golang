package organization

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Organization struct {
	ID       string    `json:"id" gorm:"primaryKey;unique;"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (org *Organization) BeforeCreate(scope *gorm.Scope) error {
	(*org).ID = uuid.New().String()
	return nil
}
