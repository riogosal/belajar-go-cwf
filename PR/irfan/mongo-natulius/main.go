package main

import (
	"app-api-natulius/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection
var ctx = context.Background()

func main() {
	var nautilus models.RawMaterials
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	c, _ := context.WithTimeout(ctx, 10*time.Second)

	defer client.Disconnect(c)
	err = client.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("natulius").Collection("raw_material_lots")
	id := "641cfe6c88032b19610a6499"
	dataIdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": dataIdObject}
	err = collection.FindOne(ctx, filter).Decode(&nautilus)
	if err != nil {
		panic(err)
	}

	nautilus.CetakIkan()
}
