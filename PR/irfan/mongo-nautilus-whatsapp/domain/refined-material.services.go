package domain

import (
	"app-api-natulius/models"
	"context"
)

type RefinedMaterialServices interface {
	GetDataByDate(context.Context, int64, int64) ([]*models.RefinedMaterialLots, error)
}
