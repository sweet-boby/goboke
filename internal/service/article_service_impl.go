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
	list := []model.Article{}
	arts, err := s.articleRepo.FindAll()

	if err != nil {
		return nil, err
	}

	for _, v := range arts {
		if v.Status != model.StatusDeleted {
			list = append(list, v)
		}
	}

	return list, nil

}

func (s *ArticleService) GetArticlesWithAdmin() ([]model.Article, error) {
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

	if req.Status == nil {
		s := model.StatusDraft
		req.Status = &s
	}

	now := time.Now()

	article := model.Article{
		Title:     *req.Title,
		Content:   *req.Content,
		Author:    *req.Author,
		UserID:    *req.UserID,
		Tag:       *req.Tag,
		Status:    *req.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.articleRepo.Create(article)
}

func (s *ArticleService) UpdateArticle(id int, req dto.UpdateArticleRequest) (*model.Article, error) {
	if req.Title == nil {
		return nil, errors.New("title is required")
	}

	if req.Content == nil {
		return nil, errors.New("content is required")
	}

	if req.Author == nil {
		return nil, errors.New("author is required")
	}

	if req.Tag == nil {
		art, err := s.articleRepo.FindByID(id)
		if err != nil {
			return nil, err
		}
		req.Tag = &art.Tag
	}

	article := model.Article{
		Title:     *req.Title,
		Content:   *req.Content,
		Author:    *req.Author,
		Tag:       *req.Tag,
		Status:    *req.Status,
		UpdatedAt: time.Now(),
	}

	return s.articleRepo.Update(id, article)
}

func (s *ArticleService) DeleteArticle(id int, userID int, role model.UserRole) error {
	art, err := s.articleRepo.FindByID(id)

	if err != nil {
		return err
	}

	if art.UserID == userID {
		return s.articleRepo.Delete(id)
	}

	if role == model.RoleAdmin {
		return s.articleRepo.Delete(id)
	}

	if art.Status == model.StatusDeleted {
		return nil
	}

	return errors.New("delete fail")
}

func (s *ArticleService) RecoverArticle(id int) error {
	art, err := s.articleRepo.FindByID(id)

	if err != nil {
		return err
	}

	if art.Status == model.StatusDeleted {
		art.Status = model.StatusDraft
	} else {
		return errors.New("article didn't be deleted")
	}

	return nil
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
