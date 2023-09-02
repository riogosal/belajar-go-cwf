package domain

import "whatsapp-nautilus/models"

type RawMaterialLotRepository interface {
	GetData(id string) (*models.RawMaterial, error)
}
