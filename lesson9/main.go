package main

import (
	"belajar-gin/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/todos", todo.Get)
	r.POST("/todos", todo.Create)
	r.PATCH("/todos/:id", todo.Patch)
	// r.DELETE("/todos/:id", todo.Delete) not implmemented

	r.Run(":1235")
}
