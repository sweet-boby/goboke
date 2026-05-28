package dto

type CreateArticleRequest struct {
	Title   *string  `json:"title" binding:"required"`
	Content *string  `json:"content" binding:"required"`
	Author  *string  `json:"-"`
	UserID  *int     `json:"-"`
	Tag     []string `json:"tag"`
}

type UpdateArticleRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Author  string   `json:"author" binding:"required"`
	Tag     []string `json:"tag"`
}
