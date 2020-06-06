package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stationedabroad/go-gin-http/controller"
	"github.com/stationedabroad/go-gin-http/service"
)

var (
	port            string                     = ":8080"
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/save", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(port)
}
