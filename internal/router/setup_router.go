package router

import (
	"goboke/internal/handler"
	"goboke/internal/middleware"
	"goboke/internal/service"

	"github.com/gin-gonic/gin"
)

type Deps struct {
	ArticleService *service.ArticleService
}

func SetupRouter(deps Deps) *gin.Engine {
	// TODO: Create Gin router without default middleware
	// Use gin.New() instead of gin.Default()
	router := gin.New()
	// TODO: Setup custom middleware in correct order
	// 1. ErrorHandlerMiddleware (first to catch panics)
	// 2. RequestIDMiddleware
	// 3. LoggingMiddleware
	// 4. CORSMiddleware
	// 5. RateLimitMiddleware
	// 6. ContentTypeMiddleware
	router.Use(
		middleware.ErrorHandlerMiddleware(),
		middleware.RequestIDMiddleware(),
		middleware.LoggingMiddleware(),
		middleware.CORSMiddleware(),
		middleware.RateLimitMiddleware(),
		middleware.ContentTypeMiddleware(),
	)

	// TODO: Setup route groups
	// Public routes (no authentication required)
	// Protected routes (require authentication)
	articleHandler := handler.NewAricleHandler(deps.ArticleService)
	// TODO: Define routes
	// Public: GET /ping, GET /articles, GET /articles/:id
	// Protected: POST /articles, PUT /articles/:id, DELETE /articles/:id, GET /admin/stats

	router.GET("/ping", handler.Ping)
	router.GET("/articles", articleHandler.GetArticles)
	router.GET("/articles/:id", articleHandler.GetArticle)
	router.POST("/articles", articleHandler.CreateArticle)
	router.PUT("/articles/:id", articleHandler.UpdateArticle)
	router.DELETE("/articles/:id", articleHandler.DeleteArticle)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/articles", articleHandler.CreateArticle)
		api.PUT("/articles/:id", articleHandler.UpdateArticle)
		api.DELETE("/articles/:id", articleHandler.DeleteArticle)
		api.GET("/admin/stats", articleHandler.GetStats)
	}
	return router
}
