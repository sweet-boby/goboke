package repository

import (
	"errors"
	"goboke/internal/model"
	"time"
)

type MemoryArticleHistoryRepository struct {
	articleHistories []model.ArticleHistory
	nextID           int
}

func NewArticleHistoryRepository() *MemoryArticleHistoryRepository {
	return &MemoryArticleHistoryRepository{
		articleHistories: []model.ArticleHistory{},
		nextID:           1,
	}
}

func (r *MemoryArticleHistoryRepository) FindByUserID(id int) ([]model.ArticleHistory, error) {
	res := []model.ArticleHistory{}

	for _, v := range r.articleHistories {
		if v.UserID == id {
			res = append(res, v)
		}
	}

	return res, nil
}

func (r *MemoryArticleHistoryRepository) Create(articleHistory model.ArticleHistory) (*model.ArticleHistory, error) {
	articleHistory.ID = r.nextID
	articleHistory.LookAt = time.Now()
	r.nextID += 1
	r.articleHistories = append(r.articleHistories, articleHistory)
	return &articleHistory, nil
}

func (r *MemoryArticleHistoryRepository) Delete(id int) error {
	for i, v := range r.articleHistories {
		if v.ID == id {
			r.articleHistories = append(r.articleHistories[:i], r.articleHistories[i+1:]...)
			return nil
		}
	}

	return errors.New("delete fail")
}
