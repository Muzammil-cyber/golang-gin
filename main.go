package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/muzammil-cyber/golang-gin/controller"
	"github.com/muzammil-cyber/golang-gin/middleware"
	"github.com/muzammil-cyber/golang-gin/repository"
	"github.com/muzammil-cyber/golang-gin/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
	jwtService      service.JWTService         = service.NewJWTService()
	loginService    service.LoginService       = service.NewLoginService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
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
		gindump.Dump())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(200, gin.H{
				"token": token,
			})
		}
	})

	apiRoutes := server.Group("/api", middleware.JWTAuthMiddleware(jwtService))
	{
		apiRoutes.POST("/videos",
			videoController.Save)

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			videos := videoController.GetAll(ctx)
			ctx.JSON(200, videos)
		})
		apiRoutes.GET("/videos/:id", func(ctx *gin.Context) {
			video := videoController.GetByID(ctx)
			ctx.JSON(200, video)
		})
		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			updatedVideo := videoController.Update(ctx)
			ctx.JSON(200, updatedVideo)
		})
		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"message": "Video deleted successfully",
			})
		})
	}

	viewRoutes := server.Group("/")
	{
		viewRoutes.GET("/", videoController.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == "" {
		// Elastic Beanstalk sets the default port to 5000
		port = "5000"
	}

	server.Run(":" + port)
}
