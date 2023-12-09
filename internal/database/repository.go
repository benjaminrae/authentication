package database

import (
	"context"

	"github.com/benjaminrae/authentication/internal/database/models"
	"github.com/google/uuid"
)

type CRUDRepository[T models.Model] interface {
	Create(ctx context.Context, model T) (*T, error)
	FindById(ctx context.Context, id uuid.UUID) (*T, error)
	FindAll(ctx context.Context) ([]T, error)
	DeleteById(ctx context.Context) error
	UpdateById(ctx context.Context, id uuid.UUID, updated T) (*T, error)
}
