package domain

import (
	"app-api-natulius/models"
	"context"
)

type PackingService interface {
	GetDataByDate(context.Context, int64, int64) ([]*models.Packings, error)
}
