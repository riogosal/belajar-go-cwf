package controller

import (
	"eser-gin/model"

	"github.com/gin-gonic/gin"
)

func DataFilm(c *gin.Context) {
	films := model.Films

	if films == nil {
		c.JSON(201, gin.H{"message": "Tidak ada data film yang tersedia"})
		return
	}
	c.JSON(201, gin.H{"data": films})
}
