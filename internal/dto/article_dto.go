package dto

import "goboke/internal/model"

type CreateArticleRequest struct {
	Title   *string              `json:"title" binding:"required"`
	Content *string              `json:"content" binding:"required"`
	Author  *string              `json:"-"`
	UserID  *int                 `json:"-"`
	Status  *model.ArticleStatus `json:"status"`
	Tag     *[]string            `json:"tag"`
}

type UpdateArticleRequest struct {
	Title   *string              `json:"title" binding:"required"`
	Content *string              `json:"content" binding:"required"`
	Author  *string              `json:"author" binding:"required"`
	Status  *model.ArticleStatus `json:"status"`
	Tag     *[]string            `json:"tag"`
}
