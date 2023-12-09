package models

import "github.com/google/uuid"

type IdentityCredentials struct {
	Password   string
	IdentityId uuid.UUID
}
