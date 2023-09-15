package domain

import (
	"app-api-natulius/models"
	"context"
)

type InvoiceService interface {
	GetData(context.Context, *string) (*models.Invoices, error)
	CreateData(context.Context, *models.Invoices) error
	GetAll(context.Context) ([]*models.Invoices, error)
	Delete(context.Context, *string) error
}
