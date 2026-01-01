package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/muzammil-cyber/golang-gin/controller"
	"github.com/muzammil-cyber/golang-gin/middleware"
	"github.com/muzammil-cyber/golang-gin/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	// Log to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuthMiddleware("admin", "password"), gindump.Dump())

	server.POST("/videos", func(ctx *gin.Context) {
		videoController.Save(ctx)
	})

	server.GET("/videos", func(ctx *gin.Context) {
		videos := videoController.GetAll(ctx)
		ctx.JSON(200, videos)
	})

	server.Run(":8080")
}
