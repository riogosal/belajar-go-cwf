package repository

import (
	"context"
	"go-mongo/collections"
	"go-mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilmRepository interface {
	FindByID(ctx context.Context, id int) (*model.Film, error)
	FindAll(ctx context.Context) ([]*model.Film, error)
	AddFilm(ctx context.Context, film *model.Film) (*mongo.InsertOneResult, error)
	Delete(ctx context.Context, id int) error
	// Save(ctx context.Context, film *model.Film) error
	// Update(ctx context.Context, id int, film *model.Film) error
}

type FilmRepositoryImplementasi struct {
	Collection *mongo.Collection
}

func NewFilmRepository(collection *mongo.Collection) FilmRepository {
	return &FilmRepositoryImplementasi{
		Collection: collection,
	}
}

func (r *FilmRepositoryImplementasi) FindByID(ctx context.Context, id int) (*model.Film, error) {
	db := r.Collection.Database() // Menggunakan koleksi yang sudah ada
	filmCollections := db.Collection(collections.Film)

	cursor, err := filmCollections.Find(ctx, bson.M{"id": id})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var results []*model.Film
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) > 0 {
		return results[0], nil
	} else {
		return nil, mongo.ErrNoDocuments
	}
}

func (r *FilmRepositoryImplementasi) FindAll(ctx context.Context) ([]*model.Film, error) {
	db := r.Collection.Database() // Menggunakan koleksi yang sudah ada
	filmCollections := db.Collection(collections.Film)

	cursor, err := filmCollections.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var films []*model.Film
	if err := cursor.All(ctx, &films); err != nil {
		return nil, err
	}

	return films, nil
}

func (r *FilmRepositoryImplementasi) AddFilm(ctx context.Context, film *model.Film) (*mongo.InsertOneResult, error) {
	db := r.Collection.Database()
	filmCollections := db.Collection(collections.Film)

	result, err := filmCollections.InsertOne(ctx, film)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *FilmRepositoryImplementasi) Delete(ctx context.Context, id int) error {
	db := r.Collection.Database()
	filmCollections := db.Collection(collections.Film)

	cekData, err := filmCollections.DeleteOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return err
	}

	if cekData.DeletedCount == 0 {
		return nil
	}

	return nil
}
