package repositories

import (
	"app-api-natulius/constants"
	"app-api-natulius/domain"
	"app-api-natulius/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RawMaterialServiceImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewRawMaterialService(db *mongo.Database) domain.RawMaterialsService {
	return &RawMaterialServiceImpl{
		db:   db,
		coll: db.Collection(string(constants.RawMaterialLots)),
	}
}

func (u *RawMaterialServiceImpl) GetData(ctx context.Context, id string) (*models.RawMaterials, error) {
	var nautilus *models.RawMaterials
	dataIdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": dataIdObject}
	err = u.coll.FindOne(ctx, filter).Decode(&nautilus)
	if err != nil {
		panic(err)
	}

	return nautilus, err

}

func (u *RawMaterialServiceImpl) GetDataByDate(ctx context.Context, startDate int64, endDate int64) ([]*models.RawMaterials, error) {

	filter := bson.M{"created_at": bson.M{"$gte": startDate, "$lte": endDate}}
	cursor, err := u.coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var data models.RawMaterials
		if err := cursor.Decode(&data); err != nil {
			panic(err)
		}

		models.DataRawMaterials = append(models.DataRawMaterials, &data)
	}

	return models.DataRawMaterials, err
}
