package domain

import (
	"app-api-natulius/models"
	"context"
)

type RawMaterialsService interface {
	GetData(context.Context, string) (*models.RawMaterials, error)
	GetDataByDate(context.Context, int64, int64) ([]*models.RawMaterials, error)
}
