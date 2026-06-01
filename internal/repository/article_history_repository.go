package repository

import "goboke/internal/model"

type ArticleHistoryRepository interface {
	FindByUserID(id int) ([]model.ArticleHistory, error)
	Create(articleHistory model.ArticleHistory) (*model.ArticleHistory, error)
	Delete(id int) error
}
