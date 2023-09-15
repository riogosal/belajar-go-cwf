package domain

import (
	"chenwoo-simple-invoices/models"
	"context"
)

type Invoices interface {
	GetData(context.Context) ([]*models.Invoices, error)
	GetDataID(context.Context, string) (*models.Invoices, error)
	CreateData(context.Context, *models.Invoices) error
	DeleteData(context.Context, string) error
	AddDataConten(context.Context, string, *models.Content) error
	DeleteDataContent(context.Context, string, string) error
}
