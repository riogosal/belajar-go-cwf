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

	// if err := c.ShouldBindJSON(&film); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest,
	// 		gin.H{
	// 			"error":   "VALIDATEERR-1",
	// 			"message": "Invalid inputs. Please check your inputs"})
	// 	return
	// }

	model.Films = append(model.Films, film)

	c.JSON(201, gin.H{"message": "Film berhasil ditambahkan", "data": film})
}
