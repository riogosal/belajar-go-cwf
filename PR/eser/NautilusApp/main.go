package main

import (
	"context"
	"fmt"
	"log"
	"nautilus-app/domain"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	return client.Database("nautilus").Collection(coll)
}

func main() {
	ctx := context.Background()
	r := gin.Default()
	client := MongoConn()
	defer client.Disconnect(context.Background())

	collection := MongoCollection("raw_material_lots", client)

	r.GET("/get-all-data", func(c *gin.Context) {
		var results []domain.RawMaterial
		filter := bson.M{"type.whole": true}

		// cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetLimit(10))
		cursor, err := collection.Find(ctx, filter)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var result domain.RawMaterial
			if err := cursor.Decode(&result); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			results = append(results, result)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/get-by-id/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		dataId, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(500, gin.H{"error": "Invalid ID format"})
			return
		}
		filter := bson.M{"_id": dataId}

		// var result []domain.GroupedContents
		var result domain.RawMaterial
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			c.JSON(404, gin.H{"error": "Data not found"})
			return
		}

		c.JSON(200, result.GroupedContents)
		result.CetakData()
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
