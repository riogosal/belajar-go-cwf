package controller

import (
	"eser-gin/model"

	"github.com/gin-gonic/gin"
)

func AddFilm(c *gin.Context) {
	var film model.Film

	if err := c.BindJSON(&film); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	model.Films = append(model.Films, film)

	c.JSON(201, gin.H{"message": "Film berhasil ditambahkan", "data": film})
}
