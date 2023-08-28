package controller

import (
	"go-mongo/collections"
	"go-mongo/koneksi"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func HapusDataByID(c *gin.Context) {

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "Back Home"})
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

	deleteResult, err := filmCollections.DeleteOne(koneksi.Ctx, bson.M{"id": id})

	if err != nil {
		panic(err)
	}

	if deleteResult.DeletedCount == 0 {
		c.JSON(404, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(200, gin.H{"message": "Data berhasil dihapus"})
}
