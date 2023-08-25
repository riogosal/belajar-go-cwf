package route

import (
	"api-movie-gin/model"

	"github.com/gin-gonic/gin"
)

func SearchHandler(ctx *gin.Context) {
	title := ctx.Query("title")

	for _, film := range model.Film {
		if film.Title == title {

			ctx.JSON(200, film)
		}
	}

}
