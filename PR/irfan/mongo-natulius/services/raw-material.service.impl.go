package services

import (
	"app-api-natulius/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RawMaterialServiceImpl struct {
	raw_material_lots *mongo.Collection
	ctx               context.Context
}

func NewRawMaterialService(RawMaterialsCollection *mongo.Collection, ctx context.Context) RawMaterialsService {
	return &RawMaterialServiceImpl{
		raw_material_lots: RawMaterialsCollection,
		ctx:               ctx,
	}
}

func (u *RawMaterialServiceImpl) GetData(id string) (*models.RawMaterials, error) {
	var nautilus *models.RawMaterials
	dataIdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": dataIdObject}
	err = u.raw_material_lots.FindOne(u.ctx, filter).Decode(&nautilus)
	if err != nil {
		panic(err)
	}

	return nautilus, err

}
