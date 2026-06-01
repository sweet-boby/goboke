package main

import (
	"goboke/internal/model"
	"goboke/internal/repository"
	"goboke/internal/router"
	"goboke/internal/service"

	"errors"
)

func main() {
	articleRepo := repository.NewMemoryArticleRepository()
	articleService := service.NewArticleService(articleRepo)

	userRepo := repository.NewMemoryUserRepository()
	userService := service.NewUserService(userRepo)

	router := router.SetupRouter(router.Deps{
		ArticleService: articleService,
		UserService:    userService,
	})
	// TODO: Start server on port 8080
	router.Run()
}

// TODO: Implement middleware functions

// TODO: Implement route handlers

// Helper functions

// findArticleByID finds an article by ID
func findArticleByID(id int) (*model.Article, int) {
	// TODO: Implement article lookup
	// Return article pointer and index, or nil and -1 if not found
	return nil, -1
}

// validateArticle validates article data
func validateArticle(article model.Article) error {
	// TODO: Implement validation
	// Check required fields: Title, Content, Author
	if article.Title == "" {
		return errors.New("title is nil")
	}
	if article.Content == "" {
		return errors.New("content is nil")
	}
	if article.Author == "" {
		return errors.New("author is nil")
	}
	return nil
}
