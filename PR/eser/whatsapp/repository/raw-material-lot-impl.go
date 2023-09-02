package repository

import (
	"context"
	"whatsapp-nautilus/domain"
	"whatsapp-nautilus/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RawMaterialLotRepositoryImpl struct {
	rawmateriallotcollection *mongo.Collection
	ctx                      context.Context
}

func NewRawMaterialLot(rawmateriallotcollection *mongo.Collection, ctx context.Context) domain.RawMaterialLotRepository {
	return &RawMaterialLotRepositoryImpl{
		rawmateriallotcollection: rawmateriallotcollection,
		ctx:                      ctx,
	}
}

func (r *RawMaterialLotRepositoryImpl) GetData(id string) (*models.RawMaterial, error) {
	dataId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	var rawmaterial models.RawMaterial
	err = r.rawmateriallotcollection.FindOne(r.ctx, bson.M{"_id": dataId}).Decode(&rawmaterial)
	return &rawmaterial, err
}
