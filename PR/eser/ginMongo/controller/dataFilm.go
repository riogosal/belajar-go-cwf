package controller

import (
	"go-mongo/collections"
	"go-mongo/koneksi"
	"go-mongo/model"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DataFilm(c *gin.Context) {

	mongo_client, err := koneksi.InitializeMongo()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := koneksi.DisconnectMongo(); err != nil {
			panic(err)
		}
	}()

	db := mongo_client.Database(os.Getenv("DB_NAME"))
	filmCollections := db.Collection(collections.Film)

	cursor, err := filmCollections.Find(koneksi.Ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	var films []model.Film
	if err := cursor.All(koneksi.Ctx, &films); err != nil {
		panic(err)
	}

	if len(films) == 0 {
		c.JSON(404, gin.H{"message": "Tidak ada data film yang tersedia"})
		return
	}

	c.JSON(200, gin.H{"data": films})
}
