package controller

import (
	"go-mongo/collections"
	"go-mongo/koneksi"
	"go-mongo/model"
	"os"

	"github.com/gin-gonic/gin"
)

func AddFilm(c *gin.Context) {
	var film model.Film

	mongoClient, err := koneksi.InitializeMongo()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := koneksi.DisconnectMongo(); err != nil {
			panic(err)
		}
	}()

	if err := c.BindJSON(&film); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// validator := validator.New()
	// if err := validator.Struct(film); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	db := mongoClient.Database(os.Getenv("DB_NAME"))
	filmCollections := db.Collection(collections.Film)

	insertResult, err := filmCollections.InsertOne(koneksi.Ctx, film)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Film berhasil ditambahkan", "data": insertResult.InsertedID})
}
