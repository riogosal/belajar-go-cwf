package todo

import (
	"belajar-gin/domain"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Patch(ctx *gin.Context) {
	// /todos?search=buy%20milk
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	var body domain.Todo
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if body.ID != id {
		ctx.AbortWithError(400, fmt.Errorf("ID in body and path does not match"))
		return
	}

	// Look for todo with id
	for i, todo := range domain.TodosInMemory {
		if todo.ID == id {
			domain.TodosInMemory[i] = todo
			ctx.JSON(200, gin.H{"status": "success"})
			return
		}
	}

	// If for loop above does not return, then todo with id is not found
	ctx.AbortWithError(404, fmt.Errorf("todo with id %d not found", id))
}
