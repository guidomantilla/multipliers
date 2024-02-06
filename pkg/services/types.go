package services

import (
	"context"

	"multipliers/pkg/models"
)

type NumbersService interface {
	Save(ctx context.Context, number *models.Number) error
	Find(ctx context.Context, number *models.Number) (*models.Number, error)
	FindAll(ctx context.Context) ([]string, error)
}
