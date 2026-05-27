package repository

import "goboke/internal/model"

type ArticleRepository interface {
	FindAll() ([]model.Article, error)
	FindByID(id int) (*model.Article, error)
	Create(article model.Article) (*model.Article, error)
	Update(id int, article model.Article) (*model.Article, error)
	Delete(id int) error
}
