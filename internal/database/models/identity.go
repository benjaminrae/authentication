package models

import (
	"time"

	"github.com/google/uuid"
)

type Identity struct {
	*Model
	Email      string
	IsVerified bool
	Username   string
	FirstName  string
	LastName   string
	BirthDate  time.Time
	UserId     uuid.UUID
}
