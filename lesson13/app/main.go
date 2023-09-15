package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb://my-db:27017",
	))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	fmt.Println("success connect to mongodb")
}
