package controller

import (
	"go-mongo/collections"
	"go-mongo/koneksi"
	"go-mongo/model"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateFilmByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID harus berupa angka"})
		return
	}

	var updatedFilm model.Film
	if err := c.BindJSON(&updatedFilm); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mongoClient, err := koneksi.InitializeMongo()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := koneksi.DisconnectMongo(); err != nil {
			panic(err)
		}
	}()

	db := mongoClient.Database(os.Getenv("DB_NAME"))
	filmCollections := db.Collection(collections.Film)

	updateFields := make(map[string]interface{})

	if updatedFilm.Judul != "" {
		updateFields["judul"] = updatedFilm.Judul
	}

	if updatedFilm.Deskripsi != "" {
		updateFields["deskripsi"] = updatedFilm.Deskripsi
	}

	if updatedFilm.Studio != "" {
		updateFields["studio"] = updatedFilm.Studio
	}

	if updatedFilm.Rating != 0 {
		updateFields["rating"] = updatedFilm.Rating
	}

	if updatedFilm.Durasi != 0 {
		updateFields["durasi"] = updatedFilm.Durasi
	}

	if len(updatedFilm.Genres) > 0 {
		updateFields["genres"] = updatedFilm.Genres
	}

	_, err = filmCollections.UpdateOne(koneksi.Ctx, bson.M{"id": id}, bson.M{"$set": updatedFilm})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Film berhasil diperbarui"})
}
