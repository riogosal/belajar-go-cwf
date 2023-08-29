package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoURI() string {
	return fmt.Sprintf(
		"mongodb://%s", // test:test@localhost:27017
		os.Getenv("DB_URI"),
	)
}

func MgoConn() *mongo.Client {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(MongoURI()).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongo Server Successfully connected")
	return client

}

func MgoCollection(person string, client *mongo.Client) *mongo.Collection {
	db := os.Getenv("DB_NAME")
	return client.Database(db).Collection(person)
}
