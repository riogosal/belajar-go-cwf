package route

import (
	"api-movie-gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteHandler(ctx *gin.Context) {

	idStr := ctx.Param("id")
	k, _ := strconv.Atoi(idStr)

	for i, m := range model.Film {
		if m.Id == k {
			model.Film = append(model.Film[:i], model.Film[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"msg": "movie deleted successfully"})

		}
	}

}
