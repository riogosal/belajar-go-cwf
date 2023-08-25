package routes

import (
	"eser-gin/controller"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()

	router.GET("/film", controller.DataFilm)
	router.POST("/add-film", controller.AddFilm)
	router.GET("/search-genre/genre", controller.SearchFilmByGenre)
	router.GET("/getid-film/:id", controller.SearchFilmById)
	router.GET("/hapus-film/:id", controller.HapusDataByID)
	router.PATCH("/update-film/:id", controller.UpdateDataByID)

	fmt.Println("Running server")
	err := router.Run(":" + os.Getenv("APP_PORT"))

	if err != nil {
		panic(err)
	}
}
