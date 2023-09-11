package domain

import (
	"app-api-natulius/models"
	"context"
)

type ProductService interface {
	GetDataByDate(context.Context, int64, int64) ([]*models.ProductLogContainers, error)
}
