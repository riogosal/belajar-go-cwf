package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, []int{1, 2, 3, 4, 5})
	})

	router.Run(":6060")
}
