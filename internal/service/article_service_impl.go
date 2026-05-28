package service

import (
	"errors"
	"goboke/internal/dto"
	"goboke/internal/model"
	"goboke/internal/repository"
	"time"
)

type ArticleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
	}
}

func (s *ArticleService) GetArticles() ([]model.Article, error) {
	return s.articleRepo.FindAll()
}

func (s *ArticleService) GetArticle(id int) (*model.Article, error) {
	return s.articleRepo.FindByID(id)
}

func (s *ArticleService) CreateArticle(req dto.CreateArticleRequest) (*model.Article, error) {
	if req.Title == nil {
		return nil, errors.New("title is required")
	}

	if req.Content == nil {
		return nil, errors.New("content is required")
	}

	if req.Author == nil {
		return nil, errors.New("author is required")
	}

	if req.UserID == nil {
		return nil, errors.New("UserID is required")
	}

	now := time.Now()

	article := model.Article{
		Title:     *req.Title,
		Content:   *req.Content,
		Author:    *req.Author,
		UserID:    *req.UserID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.articleRepo.Create(article)
}

func (s *ArticleService) UpdateArticle(id int, req dto.UpdateArticleRequest) (*model.Article, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}

	if req.Content == "" {
		return nil, errors.New("content is required")
	}

	if req.Author == "" {
		return nil, errors.New("author is required")
	}

	article := model.Article{
		Title:     req.Title,
		Content:   req.Content,
		Author:    req.Author,
		UpdatedAt: time.Now(),
	}

	return s.articleRepo.Update(id, article)
}

func (s *ArticleService) DeleteArticle(id int) error {
	return s.articleRepo.Delete(id)
}

func (s *ArticleService) GetStats() (map[string]interface{}, error) {
	articles, err := s.articleRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_articles": len(articles),
	}, nil
}
