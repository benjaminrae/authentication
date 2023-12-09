package models

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func CreateModel() Model {
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()

	return Model{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
