package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stationedabroad/go-gin-http/controller"
	"github.com/stationedabroad/go-gin-http/middlewares"
	"github.com/stationedabroad/go-gin-http/service"
)

var (
	defaultPort            string                     = ":8080"
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOuput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOuput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(defaultPort)
}
