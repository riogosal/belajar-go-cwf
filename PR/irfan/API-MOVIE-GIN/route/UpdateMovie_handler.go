package route

import (
	"api-movie-gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(ctx *gin.Context) {
	var dataMovie model.Movie

	IdString := ctx.Param("id")

	k, _ := strconv.Atoi(IdString)

	if err := ctx.ShouldBindJSON(&dataMovie); err != nil {
		panic("error updating dataMovie !")
	}

	for i, v := range model.Film {
		if v.Id == k {
			model.Film[i] = dataMovie

			ctx.JSON(http.StatusOK, gin.H{"msg": "Update Data Movie successfully"})

		}
	}
}
