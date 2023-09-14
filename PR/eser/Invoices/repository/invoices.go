package repository

import (
	"chenwoo-simple-invoices/constants"
	"chenwoo-simple-invoices/domain"
	"chenwoo-simple-invoices/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoicesImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewInvoices(db *mongo.Database) domain.Invoices {
	return &InvoicesImpl{
		db:   db,
		coll: db.Collection(string(constants.InvoicesCollectionName)),
	}
}

func (p *InvoicesImpl) GetData(ctx context.Context) ([]*models.Invoices, error) {
	filter := bson.M{}

	var invoices []*models.Invoices

	cursor, err := p.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var dataInvoices *models.Invoices

		if err := cursor.Decode(&dataInvoices); err != nil {
			return nil, err
		}

		invoices = append(invoices, dataInvoices)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (p *InvoicesImpl) GetDataID(ctx context.Context, id string) (*models.Invoices, error) {
	dataId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	var dataInvoices *models.Invoices

	err = p.coll.FindOne(ctx, bson.M{"_id": dataId}).Decode(&dataInvoices)

	return dataInvoices, err
}

func (p *InvoicesImpl) CreateData(ctx context.Context, invoices *models.Invoices) error {
	for i := range invoices.Contents {
		newUUID := uuid.New()
		uuidString := newUUID.String()

		invoices.Contents[i].UUID = uuidString
	}

	newUUID := uuid.New()
	uuidString := newUUID.String()
	invoices.UUID = uuidString

	result, err := p.coll.InsertOne(ctx, invoices)
	if err != nil {
		return err
	}

	var totalCount int64
	var totalPrice int64

	for _, v := range invoices.Contents {
		totalCount += v.Count
		totalPrice += v.Price
	}

	update := bson.M{
		"$inc": bson.M{
			"total_count": totalCount,
			"total_price": totalPrice,
		},
	}
	_, err = p.coll.UpdateOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}, update)

	if err != nil {
		return err
	}

	return nil
}

func (p *InvoicesImpl) DeleteData(ctx context.Context, id string) error {
	dataID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	cekData, err := p.coll.DeleteOne(ctx, bson.M{"_id": dataID})

	if err != nil {
		return nil
	}

	if cekData.DeletedCount == 0 {
		return errors.New("Data Tidak Ditemukan")
	}
	return nil
}
