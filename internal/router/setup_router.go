package router

import (
	"goboke/internal/handler"
	"goboke/internal/middleware"
	"goboke/internal/model"
	"goboke/internal/service"

	"github.com/gin-gonic/gin"
)

type Deps struct {
	ArticleService *service.ArticleService
	UserService    *service.UserService
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

	userHandler := handler.NewUserHandler(deps.UserService)

	api := router.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.GET("/ping", handler.Ping)
		api.GET("/articles", articleHandler.GetArticles)
		api.GET("/articles/:id", articleHandler.GetArticle)
		api.POST("/refresh", userHandler.RefreshToken)
	}

	auth := api.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/articles", articleHandler.CreateArticle)
		auth.DELETE("/articles/:id", articleHandler.DeleteArticle)
		auth.PUT("/articles/:id", articleHandler.UpdateArticle)
		auth.POST("/logout", userHandler.Logout)

		auth.GET("/profile", userHandler.GetUserProfile)
		auth.PUT("/profile", userHandler.UpdateUserProfile)
	}

	admin := api.Group("/admin")

	admin.Use(middleware.AuthMiddleware(),
		middleware.RequireRole(string(model.RoleAdmin)))
	{
		admin.GET("/users", userHandler.GetUsers)
		admin.PUT("/users/:id/role", userHandler.UpdateUserRole)
		admin.PUT("/article/recover/:id", articleHandler.RecoverArticle)
		admin.GET("/articles", articleHandler.GetArticlesWithAdmin)
	}

	return router
}
