package controller

import (
	"eser-gin/model"

	"github.com/gin-gonic/gin"
)

func SearchFilmByGenre(c *gin.Context) {
	genreType := c.Query("genre_type")

	var genreFilm []model.Film

	for _, film := range model.Films {
		genres := film.Genres
		for _, genre := range genres {
			if string(genre) == genreType {
				genreFilm = append(genreFilm, film)
				break
			}
		}
	}

	if len(genreFilm) == 0 {
		c.JSON(404, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(201, gin.H{"data": genreFilm})
}
