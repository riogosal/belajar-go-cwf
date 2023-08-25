package controller

import (
	"eser-gin/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateDataByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}

	var updateData model.Film
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var filmToUpdate *model.Film
	for i := range model.Films {
		if model.Films[i].ID == id {
			filmToUpdate = &model.Films[i]
			break
		}
	}

	if filmToUpdate != nil {

		if updateData.Judul != "" {
			filmToUpdate.Judul = updateData.Judul
		}

		if updateData.Deskripsi != "" {
			filmToUpdate.Deskripsi = updateData.Deskripsi
		}

		if updateData.Studio != "" {
			filmToUpdate.Studio = updateData.Studio
		}

		if updateData.Rating != 0 {
			filmToUpdate.Rating = updateData.Rating
		}

		if updateData.Durasi != 0 {
			filmToUpdate.Durasi = updateData.Durasi
		}

		if len(updateData.Genres) > 0 {
			filmToUpdate.Genres = updateData.Genres
		}

		c.JSON(200, gin.H{"status": "sukses", "message": fmt.Sprintf("Film dengan ID %d telah diperbarui", id)})
	} else {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Film dengan ID %d tidak ditemukan", id)})
	}
}
