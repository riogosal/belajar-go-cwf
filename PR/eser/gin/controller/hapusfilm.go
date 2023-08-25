package controller

import (
	"eser-gin/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HapusDataByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}

	var index int = -1

	for i, film := range model.Films {
		if film.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		model.Films = append(model.Films[:index], model.Films[index+1:]...)
		c.JSON(200, gin.H{"status": "sukses", "message": fmt.Sprintf("Film dengan ID %d telah dihapus", id)})
	} else {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Film dengan ID %d tidak ditemukan", id)})
	}
}
