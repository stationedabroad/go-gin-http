package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stationedabroad/go-gin-http/controller"
	"github.com/stationedabroad/go-gin-http/middlewares"
	"github.com/stationedabroad/go-gin-http/service"
)

var (
	port            string                     = ":8080"
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(port)
}
