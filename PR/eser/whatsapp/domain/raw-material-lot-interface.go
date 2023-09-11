package domain

import (
	"context"
	"whatsapp-nautilus/models"
)

// GetManyData(ctx context.Context, filter interface{}) ([]models.RawMaterial, error)
type RawMaterialLotRepository interface {
	GetData(context.Context, string) (*models.RawMaterial, error)
	GetDataByDate(context.Context, int64, int64) ([]*models.RawMaterial, error)
}
