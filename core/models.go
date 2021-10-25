package core

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created"`
	ModifiedAt time.Time `json:"modified"`
}
