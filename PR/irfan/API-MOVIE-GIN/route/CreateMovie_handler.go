package route

import (
	"api-movie-gin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler(ctx *gin.Context) {
	var dataMovie model.Movie

	if err := ctx.ShouldBindJSON(&dataMovie); err != nil {
		log.Fatal("CreateHandler error binding json: ", err)
	}

	model.Film = append(model.Film, dataMovie)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Create Data Movie successfully"})

}
