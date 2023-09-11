package repository

import (
	"context"
	"whatsapp-nautilus/constants"
	"whatsapp-nautilus/domain"
	"whatsapp-nautilus/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PackingRepositoryImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewPacking(db *mongo.Database) domain.PackingRepository {
	return &PackingRepositoryImpl{
		db:   db,
		coll: db.Collection(string(constants.PackingCollectionName)),
	}
}

func (p *PackingRepositoryImpl) GetDataPackingByDate(ctx context.Context, startOfDayInt int64, endOfDayInt int64) ([]*models.Packing, error) {
	filter := bson.M{
		"created_at": bson.M{
			"$gte": startOfDayInt,
			"$lte": endOfDayInt,
		},
	}

	cursor, err := p.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var packing models.Packing
		if err := cursor.Decode(&packing); err != nil {
			return nil, err
			// panic(err)
		}
		models.Packings = append(models.Packings, &packing)
	}
	return models.Packings, nil
}
