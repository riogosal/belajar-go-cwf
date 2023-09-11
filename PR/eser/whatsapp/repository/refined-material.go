package repository

import (
	"context"
	"whatsapp-nautilus/constants"
	"whatsapp-nautilus/domain"
	"whatsapp-nautilus/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefinedmaterialRepositoryImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewRefinedMaterial(db *mongo.Database) domain.RefinedmaterialRepository {
	return &RefinedmaterialRepositoryImpl{
		db:   db,
		coll: db.Collection(string(constants.RefinedMaterialLotsCollectionName)),
	}
}

func (r *RefinedmaterialRepositoryImpl) GetDataRefinedByDate(ctx context.Context, startOfDayInt int64, endOfDayInt int64) ([]*models.RefinedMaterial, error) {
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
		var refinedMaterial models.RefinedMaterial
		if err := cursor.Decode(&refinedMaterial); err != nil {
			return nil, err
		}
		models.RefinedResults = append(models.RefinedResults, &refinedMaterial)
	}

	return models.RefinedResults, nil
}
