package todo

import (
	"belajar-gin/domain"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var todo domain.Todo

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	domain.TodosInMemory = append(domain.TodosInMemory, todo)

	ctx.JSON(201, gin.H{"status": "success"})
}
