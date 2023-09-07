package domain

import (
	"context"
	"whatsapp-nautilus/models"
)

type ProductLogContainersRepository interface {
	GetDataProductLogByDate(context.Context, int64, int64) ([]*models.ProductLogContainer, error)
}
