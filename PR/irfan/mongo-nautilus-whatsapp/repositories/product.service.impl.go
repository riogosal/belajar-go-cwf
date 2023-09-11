package repositories

import (
	"app-api-natulius/constants"
	"app-api-natulius/domain"
	"app-api-natulius/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewProductServiceImpl(db *mongo.Database) domain.ProductService {
	return &ProductServiceImpl{
		db:   db,
		coll: db.Collection(string(constants.ProductLogContainers)),
	}
}

func (u *ProductServiceImpl) GetDataByDate(ctx context.Context, startDate int64, endDate int64) ([]*models.ProductLogContainers, error) {
	filter := bson.M{"created_at": bson.M{"$gte": startDate, "$lte": endDate}}
	cursor, err := u.coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var data models.ProductLogContainers
		if err := cursor.Decode(&data); err != nil {
			panic(err)
		}

		models.DataProducts = append(models.DataProducts, &data)
	}

	return models.DataProducts, nil

}
