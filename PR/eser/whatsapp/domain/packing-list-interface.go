package domain

import (
	"context"
	"whatsapp-nautilus/models"
)

type PackingRepository interface {
	GetDataPackingByDate(context.Context, int64, int64) ([]*models.Packing, error)
}
