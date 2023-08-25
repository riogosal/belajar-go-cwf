package controller

import (
	"eser-gin/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchFilmById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "Sepertinya Anda Nyasar, Yuk balik sebelum gelap"})
		return
	}

	var cariIdFilm *model.Film
	for _, film := range model.Films {
		if film.ID == id {
			cariIdFilm = &film
			break
		}
	}

	if cariIdFilm != nil {
		c.JSON(200, gin.H{"status": "sukses", "data": cariIdFilm})
	} else {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Film dengan ID %d tidak ditemukan", id)})
	}
}
