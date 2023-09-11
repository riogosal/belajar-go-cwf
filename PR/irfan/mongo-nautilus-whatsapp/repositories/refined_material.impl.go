package repositories

import (
	"app-api-natulius/constants"
	"app-api-natulius/domain"
	"app-api-natulius/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefinedMaterialLotsImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewRefinedMaterialLots(db *mongo.Database) domain.RefinedMaterialServices {
	return &RefinedMaterialLotsImpl{
		db:   db,
		coll: db.Collection(string(constants.RefinedMaterialLots)),
	}
}

func (r *RefinedMaterialLotsImpl) GetDataByDate(ctx context.Context, startDate int64, EndDate int64) ([]*models.RefinedMaterialLots, error) {
	filter := bson.M{"created_at": bson.M{"$gte": startDate, "$lte": EndDate}}
	cursor, err := r.coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var data models.RefinedMaterialLots
		if err := cursor.Decode(&data); err != nil {
			panic(err)
		}
		models.DataRefinedMaterials = append(models.DataRefinedMaterials, &data)
	}

	return models.DataRefinedMaterials, err
}
