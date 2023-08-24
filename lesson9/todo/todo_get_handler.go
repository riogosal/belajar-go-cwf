package todo

import (
	"belajar-gin/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	search := ctx.Query("search") // extract "value" from ?search=value from URL
	fmt.Println("I'm trying to search for", search)

	ctx.JSON(200, domain.TodosInMemory)
}
