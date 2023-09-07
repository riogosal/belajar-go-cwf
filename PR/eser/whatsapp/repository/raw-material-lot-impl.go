package repository

import (
	"context"
	"whatsapp-nautilus/constants"
	"whatsapp-nautilus/domain"
	"whatsapp-nautilus/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RawMaterialLotRepositoryImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewRawMaterialLot(db *mongo.Database) domain.RawMaterialLotRepository {
	return &RawMaterialLotRepositoryImpl{
		db:   db,
		coll: db.Collection(string(constants.RawMaterialLotsCollectionName)),
	}
}

func (r *RawMaterialLotRepositoryImpl) GetData(ctx context.Context, id string) (*models.RawMaterial, error) {
	dataId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	var rawmaterial models.RawMaterial
	err = r.coll.FindOne(ctx, bson.M{"_id": dataId}).Decode(&rawmaterial)
	return &rawmaterial, err
}

func (r *RawMaterialLotRepositoryImpl) GetDataByDate(ctx context.Context, startOfDayInt int64, endOfDayInt int64) ([]*models.RawMaterial, error) {
	filter := bson.M{
		"created_at": bson.M{
			"$gte": startOfDayInt,
			"$lte": endOfDayInt,
		},
	}
	cursor, err := r.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var rawmaterial models.RawMaterial
		if err := cursor.Decode(&rawmaterial); err != nil {
			return nil, err
		}
		models.Results = append(models.Results, &rawmaterial)
	}

	return models.Results, nil
}
