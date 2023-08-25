package main

import (
	"belajar-mongo/collections"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID       int     `bson:"id"`
	Name     string  `bson:"name"`
	Password *string `bson:"password"`
}

func mongoURI() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s", // test:test@localhost:27017
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_URI"),
	)
}

func initializeMongo() (*mongo.Client, error) {
	ctx := context.Background()
	mongo_client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI(),
	).SetReadPreference(readpref.Primary()))
	if err != nil {
		return nil, err
	}

	return mongo_client, nil
}

func main() {
	godotenv.Load()
	ctx := context.Background()

	mongo_client, err := initializeMongo()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := mongo_client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := mongo_client.Database(os.Getenv("DB_NAME"))
	userCollection := db.Collection(collections.User)

	cursor, err := userCollection.Find(ctx, bson.D{
		{Key: "id", Value: 1},
	})
	if err != nil {
		panic(err)
	}

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		panic(err)
	}

	fmt.Println(users) // [{1 Rio "password"}, {2 Irfan nil}]
}
