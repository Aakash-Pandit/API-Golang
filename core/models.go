package core

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID       uuid.UUID `json:"id"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}
