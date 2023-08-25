package route

import (
	"api-movie-gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	k, err := strconv.Atoi(id)

	if err != nil {
		panic("error converting")
	}

	for _, v := range model.Film {
		if v.Id == k {
			ctx.JSON(http.StatusOK, v)
		}
	}
}
