package core

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("ID", uuid)
}
