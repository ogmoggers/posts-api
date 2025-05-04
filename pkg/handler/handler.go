package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"social-network-api/docs"
	"social-network-api/pkg/middleware"
	"social-network-api/pkg/service"
)


type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = "Posts API"
	docs.SwaggerInfo.Description = "API for managing posts"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8092"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret" // Default secret for development
	}

	api := router.Group("/api", middleware.JWTAuth(jwtSecret))
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getPosts)
			posts.GET("/:post_id", h.getPostById)
			posts.PATCH("/:post_id", h.updatePost)
			posts.DELETE("/:post_id", h.deletePost)
		}

		users := api.Group("/users")
		{
			users.GET("/:user_id/posts", h.getUsersPosts)
		}
	}
	return router
}
