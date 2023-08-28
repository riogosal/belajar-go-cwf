package controller

import (
	"fmt"
	"go-mongo/collections"
	"go-mongo/koneksi"
	"go-mongo/model"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchFilmById(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "Sepertinya Anda Nyasar"})
		return
	}

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

	cursor, err := filmCollections.Find(koneksi.Ctx, bson.M{"id": id})

	if err != nil {
		panic(err)
	}

	var results []model.Film
	if err := cursor.All(koneksi.Ctx, &results); err != nil {
		panic(err)
	}

	// for _, film := range results {
	// 	fmt.Printf("Data yang ditemukan:\nID: %d\nJudul: %s\nDeskripsi: %s\nStudio: %s\nRating: %.2f\nDurasi: %d\nGenres: %v\n",
	// 		film.ID, film.Judul, film.Deskripsi, film.Studio, film.Rating, film.Durasi, film.Genres)
	// }

	if results != nil {
		c.JSON(200, gin.H{"status": "sukses", "data": results})
	} else {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Film dengan ID %d tidak ditemukan", id)})
	}

}
