package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/muzammil-cyber/golang-gin/controller"
	"github.com/muzammil-cyber/golang-gin/dto"
	"github.com/muzammil-cyber/golang-gin/middleware"
	"github.com/muzammil-cyber/golang-gin/repository"
	"github.com/muzammil-cyber/golang-gin/service"
	gindump "github.com/tpkeeper/gin-dump"

	_ "github.com/muzammil-cyber/golang-gin/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/muzammil-cyber/golang-gin/graph"
	"github.com/vektah/gqlparser/v2/ast"
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

// @title Video Management API
// @version 1.0
// @description A RESTful API for managing video content with user authentication. This API allows you to create, read, update, and delete video entries along with author information. All video endpoints require JWT authentication.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support Team
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:5000
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token. Example: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(),
		gindump.Dump())

	// server.Static("/static", "./templates/static")
	// server.LoadHTMLGlob("templates/*.html")

	// swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public API routes (no JWT required)
	server.POST("/auth/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(200, dto.LoginResponse{Token: token})
		}
	})

	// Protected API routes (JWT required)
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
				ctx.JSON(500, dto.ErrorResponse{
					Error: err.Error(),
				})
				return
			}
			ctx.JSON(200, dto.MessageResponse{
				Message: "Video deleted successfully",
			})
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/", videoController.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == "" {
		// Elastic Beanstalk sets the default port to 5000
		port = "5000"
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			VideoService: videoService,
			JWTService:   jwtService,
		},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	server.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// GraphQL endpoint for mutations
	server.POST("/query", func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})

	// GraphQL queries (GET) can be public or protected based on your needs
	server.GET("/query", func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	server.Run(":" + port)
}
