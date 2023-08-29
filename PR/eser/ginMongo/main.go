package main

import (
	"go-mongo/controller"
	"go-mongo/db"
	"go-mongo/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// godotenv.Load()
// fmt.Println(os.Getenv("DB_NAME"))
// routes.SetupRoutes()

func main() {
	godotenv.Load()
	r := gin.Default()

	// Inisialisasi koneksi MongoDB sekali
	mongoClient := db.MongoConn()
	collection := db.MongoCollection("film", mongoClient)

	// Inisialisasi repository dan controller
	filmRepository := repository.NewFilmRepository(collection)
	filmController := controller.NewFilmController(filmRepository)

	// Definisikan rute-rute di sini
	r.GET("/getid-film/:id", filmController.GetFilmByID)
	r.GET("/film", filmController.GetDataFilm)
	r.POST("/addfilm", filmController.AddFilm)
	r.DELETE("/hapus-film/:id", filmController.Delete)

	r.Run(":8081")
}
