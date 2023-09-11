package repositories

import (
	"app-api-natulius/constants"
	"app-api-natulius/domain"
	"app-api-natulius/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PackingServiceImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewPackingServiceImpl(db *mongo.Database) domain.PackingService {
	return &PackingServiceImpl{
		db:   db,
		coll: db.Collection(string(constants.Packings)),
	}
}

func (s *PackingServiceImpl) GetDataByDate(ctx context.Context, startDate int64, endDate int64) ([]*models.Packings, error) {
	filter := bson.M{"created_at": bson.M{"$gte": startDate, "$lte": endDate}}
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	for cursor.Next(ctx) {
		var data models.Packings
		if err := cursor.Decode(&data); err != nil {
			panic(err)
		}
		models.DataPackings = append(models.DataPackings, &data)
	}

	return models.DataPackings, err
}
