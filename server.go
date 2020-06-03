package main

import "github.com/gin-gonic/gin"

var port string = ":8080"

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server started OK!",
		})
	})

	server.Run(port)
}
