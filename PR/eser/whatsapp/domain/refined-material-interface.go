package domain

import (
	"context"
	"whatsapp-nautilus/models"
)

type RefinedmaterialRepository interface {
	GetDataRefinedByDate(context.Context, int64, int64) ([]*models.RefinedMaterial, error)
}
