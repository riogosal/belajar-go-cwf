package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client

// var Ctx = context.Background()

func MongoConn() *mongo.Client {
	clientOptions := options.Client().
		ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Koneksi gagal:", err)
	}
	fmt.Println("Koneksi berhasil")
	return client
}

func MongoCollection(coll string, client *mongo.Client) *mongo.Collection {
	return client.Database("mydb").Collection(coll)
}
