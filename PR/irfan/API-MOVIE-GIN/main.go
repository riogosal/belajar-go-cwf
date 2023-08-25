package main

import (
	"api-movie-gin/model"
	"api-movie-gin/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

const APP_PORT = ":8000"

func main() {
	dataLabels := model.Film
	for _, label := range dataLabels {
		fmt.Println(label.Label)
		if label.Label[0] == "Romantic" {
			fmt.Println("data ada")
		}
	}

	router := gin.Default()
	router.GET("/film", route.GetHandler)
	router.GET("/film/:id", route.ShowHandler)
	router.GET("/filmSearch", route.SearchHandler)
	router.POST("/film", route.CreateHandler)
	router.PUT("/film/:id", route.UpdateHandler)
	router.DELETE("/film/:id", route.DeleteHandler)

	fmt.Println("Starting Server with port " + APP_PORT)
	router.Run(APP_PORT)

}
