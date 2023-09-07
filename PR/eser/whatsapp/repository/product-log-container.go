package repository

import (
	"context"
	"whatsapp-nautilus/constants"
	"whatsapp-nautilus/domain"
	"whatsapp-nautilus/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductLogContainersRepositoryImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewProductLogContainers(db *mongo.Database) domain.ProductLogContainersRepository {
	return &ProductLogContainersRepositoryImpl{
		db:   db,
		coll: db.Collection(string(constants.ProductLogContainersCollectionName)),
	}
}

func (p *ProductLogContainersRepositoryImpl) GetDataProductLogByDate(ctx context.Context, startOfDayInt int64, endOfDayInt int64) ([]*models.ProductLogContainer, error) {
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
		var productLogContainer models.ProductLogContainer
		if err := cursor.Decode(&productLogContainer); err != nil {
			return nil, err
			// panic(err)
		}
		models.ProductLogResult = append(models.ProductLogResult, &productLogContainer)
	}
	return models.ProductLogResult, nil
}
