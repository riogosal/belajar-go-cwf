package routes

import (
	"fmt"
	"go-mongo/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	router.GET("/film", controller.GetDataFilm)
	router.GET("/getid-film/:id", controller.SearchFilmById)
	router.POST("/add-film", controller.AddFilm)
	router.DELETE("/hapus-film/:id", controller.HapusDataByID)
	router.GET("/search-genre/genre", controller.SearchFilmByGenre)
	router.PATCH("/update-film/:id", controller.UpdateFilmByID)

	fmt.Println("Running server")
	err := router.Run(":" + os.Getenv("APP_PORT"))

	if err != nil {
		panic(err)
	}
}
