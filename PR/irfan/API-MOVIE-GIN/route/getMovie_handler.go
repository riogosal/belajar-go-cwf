package route

import (
	"api-movie-gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.Film)
}
