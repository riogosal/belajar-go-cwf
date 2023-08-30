package services

import "app-api-natulius/models"

type RawMaterialsService interface {
	GetData(string) (*models.RawMaterials, error)
}
