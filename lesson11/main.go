package main

import (
	"interfaces/film"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// mongo_film_repo := film.NewMongoFilmRepository()
	mysql_film_repo := film.NewMySQLFilmRepository()

	filmHandler := film.GinFilmHandler{
		Repo: &mysql_film_repo,
	}

	r.GET("/films", filmHandler.GetAllFilms)

}
