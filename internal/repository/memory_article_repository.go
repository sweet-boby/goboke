package repository

import (
	"errors"
	"goboke/internal/model"
	"time"
)

type MemoryArticleRepository struct {
	articles []model.Article
	nextID   int
}

func NewMemoryArticleRepository() *MemoryArticleRepository {
	// In-memory storage
	var articles = []model.Article{
		{ID: 1, Title: "Getting Started with Go", Content: "Go is a programming language...", Author: "John Doe", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Title: "Web Development with Gin", Content: "Gin is a web framework...", Author: "Jane Smith", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	var nextID = 3

	return &MemoryArticleRepository{
		articles: articles,
		nextID:   nextID,
	}
}
func (r *MemoryArticleRepository) FindAll() ([]model.Article, error) {
	return r.articles, nil
}
func (r *MemoryArticleRepository) FindByID(id int) (*model.Article, error) {

	for i, art := range r.articles {
		if art.ID == id {

			return &r.articles[i], nil
		}
	}
	return nil, errors.New("not found article")
}

func (r *MemoryArticleRepository) Create(article model.Article) (*model.Article, error) {
	article.ID = r.nextID
	r.nextID += 1
	r.articles = append(r.articles, article)

	return &article, nil
}

func (r *MemoryArticleRepository) Update(id int, article model.Article) (*model.Article, error) {
	for index, art := range r.articles {
		if art.ID == id {
			oldArt := &r.articles[index]
			oldArt.Title = article.Title
			oldArt.Content = article.Content
			oldArt.Author = article.Author
			oldArt.Tag = article.Tag
			oldArt.Status = article.Status
			oldArt.UpdatedAt = time.Now()
			return oldArt, nil
		}
	}
	return nil, errors.New("update fail")
}

func (r *MemoryArticleRepository) Delete(id int) error {
	for index, art := range r.articles {
		if art.ID == id {
			r.articles[index].Status = model.StatusDeleted
			return nil
		}
	}
	return errors.New("delete article fail")
}
