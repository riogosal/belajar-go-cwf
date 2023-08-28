package controller

import (
	"fmt"
	"go-mongo/collections"
	"go-mongo/koneksi"
	"go-mongo/model"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchFilmByGenre(c *gin.Context) {
	genreType := c.Query("genre_type")

	var genreFilm []model.Film

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

	cursor, err := filmCollections.Find(koneksi.Ctx, bson.M{"genres": genreType})
	if err != nil {
		panic(err)
	}

	if err := cursor.All(koneksi.Ctx, &genreFilm); err != nil {
		panic(err)
	}

	if len(genreFilm) > 0 {
		c.JSON(200, gin.H{"status": "sukses", "data": genreFilm})
	} else {
		c.JSON(404, gin.H{"status": "not found", "message": fmt.Sprintf("Film Dengan Genre %s Tidak Ditemukan", genreType)})
	}
}
