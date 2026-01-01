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

	server.Static("/static", "./templates/static")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Logger(),
		middleware.BasicAuthMiddleware("admin", "password"),
		gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.POST("/videos",
			videoController.Save)

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			videos := videoController.GetAll(ctx)
			ctx.JSON(200, videos)
		})
	}

	viewRoutes := server.Group("/")
	{
		viewRoutes.GET("/", videoController.ShowAll)
	}

	server.Run(":8080")
}
