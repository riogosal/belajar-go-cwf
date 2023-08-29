package route

import (
	"api-movie-gin/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateHandler(ctx *gin.Context) {
	var dataMovie model.Movie

	if err := ctx.BindJSON(&dataMovie); err != nil {

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			ctx.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	model.Film = append(model.Film, dataMovie)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Create Data Movie successfully"})

}
