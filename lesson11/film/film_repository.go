package film

import (
	"interfaces/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type SQLFilmRepository struct {
	db *gorm.DB
}

func NewMySQLFilmRepository(db *gorm.DB) domain.FilmRepository {
	return &SQLFilmRepository{
		db: db,
	}
}

type MongoFilmRepsoitory struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewMongoFilmRepository(db *mongo.Database) domain.FilmRepository {
	return &MongoFilmRepsoitory{
		db:   db,
		coll: db.Collection("films"),
	}
}

// func
