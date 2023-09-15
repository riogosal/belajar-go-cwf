package repositories

import (
	"app-api-natulius/constants"
	"app-api-natulius/domain"
	"app-api-natulius/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceSerciceImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewInvoiceSerciceImpl(db *mongo.Database) domain.InvoiceService {
	return &InvoiceSerciceImpl{
		db:   db,
		coll: db.Collection(string(constants.Invoice)),
	}
}

func (s *InvoiceSerciceImpl) GetData(ctx context.Context, id *string) (*models.Invoices, error) {
	var invoice *models.Invoices
	query := bson.D{bson.E{Key: "uuid", Value: id}}
	err := s.coll.FindOne(ctx, query).Decode(&invoice)
	return invoice, err
}

func (s *InvoiceSerciceImpl) CreateData(ctx context.Context, invoice *models.Invoices) error {
	var totPrice, totCount int64
	Uuid := uuid.New()
	invoice.UUID = Uuid.String()
	for i, _ := range invoice.Contents {
		UuidContents := uuid.New()
		invoice.Contents[i].Uuid_contents = UuidContents.String()
	}

	result, err := s.coll.InsertOne(ctx, invoice)
	if err != nil {
		panic(err)

	}

	for _, v := range invoice.Contents {
		totCount += v.Count
		totPrice += v.Price
	}
	updateContents := bson.M{"$set": bson.M{"total_price": totPrice, "total_count": totCount}}
	_, err = s.coll.UpdateOne(ctx, bson.D{{Key: "_id", Value: result.InsertedID}}, updateContents)

	return err

}

func (s *InvoiceSerciceImpl) GetAll(ctx context.Context) ([]*models.Invoices, error) {
	cursor, err := s.coll.Find(ctx, bson.D{{}})

	for cursor.Next(ctx) {
		var invoice models.Invoices
		err := cursor.Decode(&invoice)
		if err != nil {
			return nil, err
		}
		models.DataInvoices = append(models.DataInvoices, &invoice)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)
	if len(models.DataInvoices) == 0 {
		return nil, errors.New("Document not found")
	}

	return models.DataInvoices, err
}

func (s *InvoiceSerciceImpl) Delete(ctx context.Context, id *string) error {
	filter := bson.D{primitive.E{Key: "uuid", Value: id}}
	result, _ := s.coll.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("Document not found")
	}

	return nil
}
