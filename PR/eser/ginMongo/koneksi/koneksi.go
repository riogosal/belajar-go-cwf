package koneksi

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// var Ctx = context.TODO()
var Ctx = context.Background()

func mongoURI() string {
	return fmt.Sprintf(
		"mongodb://%s",
		os.Getenv("DB_URI"),
	)
}

func InitializeMongo() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}

	clientOptions := options.Client().ApplyURI(mongoURI())

	newClient, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = newClient.Ping(Ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	client = newClient
	return client, nil
}

func DisconnectMongo() error {
	if client == nil {
		return nil
	}

	err := client.Disconnect(Ctx)
	if err != nil {
		return err
	}

	client = nil
	return nil
}
