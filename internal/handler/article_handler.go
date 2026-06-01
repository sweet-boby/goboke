package handler

import (
	"goboke/internal/dto"
	"goboke/internal/model"
	"goboke/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService *service.ArticleService
}

func NewAricleHandler(articleService *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
	}
}

// getStats handles GET /admin/stats - get API usage statistics (admin only)
func (h *ArticleHandler) GetStats(c *gin.Context) {
	// TODO: Check if user role is "admin"
	// TODO: Return mock statistics
	arts, err := h.articleService.GetStats()

	if err != nil {
		c.JSON(404, dto.APIResponse{
			Success: false,
		})
		return
	}

	stats := map[string]interface{}{
		"total_articles": len(arts),
		"total_requests": 0, // Could track this in middleware
		"uptime":         "24h",
	}
	role, ok := c.Get("user_role")

	if ok == false {
		c.JSON(403, gin.H{
			"success": false,
		})
		return
	}

	if role != "admin" {
		c.JSON(403, gin.H{
			"success": false,
		})
		return
	}
	// TODO: Return stats in standard format
	c.JSON(200, gin.H{
		"success": true,
		"data":    stats,
	})

	// TODO: Return stats in standard format
}

func (h *ArticleHandler) GetArticles(c *gin.Context) {
	// TODO: Implement pagination (optional)
	// TODO: Return articles in standard format
	data, _ := c.Get("request_id")
	arts, err := h.articleService.GetArticles()

	if err != nil {
		c.JSON(404, dto.APIResponse{
			Success: false,
		})
	}

	c.JSON(200, gin.H{
		"success":    true,
		"data":       arts,
		"request_id": data,
	})
}

// getArticle handles GET /articles/:id - get article by ID
func (h *ArticleHandler) GetArticle(c *gin.Context) {
	// TODO: Get article ID from URL parameter
	// TODO: Find article by ID
	// TODO: Return 404 if not found
	id := c.Param("id")
	artID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}

	art, err := h.articleService.GetArticle(artID)

	c.JSON(200, dto.APIResponse{
		Success: true,
		Data:    art,
	})

}

// createArticle handles POST /articles - create new article (protected)
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	// TODO: Parse JSON request body
	// TODO: Validate required fields
	// TODO: Add article to storage
	// TODO: Return created article
	var article dto.CreateArticleRequest
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}

	userID, ok := c.Get("userID")

	if !ok {
		c.JSON(400, dto.APIResponse{
			Success: true,
			Message: "userID not fonund",
		})
		return
	}

	username, ok := c.Get("username")

	if !ok {
		c.JSON(400, dto.APIResponse{
			Success: true,
			Message: "userName not fonund",
		})
		return
	}

	userIDInt := userID.(int)
	article.UserID = &userIDInt
	userNameStr := username.(string)
	article.Author = &userNameStr

	art, err := h.articleService.CreateArticle(article)

	if err != nil {

	}
	c.JSON(201, gin.H{
		"success": true,
		"data":    art,
	})
}

// deleteArticle handles DELETE /articles/:id - delete article (protected)
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	// TODO: Get article ID from URL parameter
	// TODO: Find and remove article
	// TODO: Return success message
	id := c.Param("id")
	artID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(404, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {

		c.JSON(404, dto.APIResponse{
			Success: false,
			Message: "not found userid",
		})
		return
	}

	role, ok := c.Get("role")

	if !ok {
		c.JSON(404, dto.APIResponse{
			Success: false,
			Message: "not found user role",
		})
		return
	}

	err = h.articleService.DeleteArticle(artID, userID.(int), model.UserRole(role.(string)))

	if err != nil {
		c.JSON(400, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, dto.APIResponse{
		Success: true,
		Message: "delete success",
	})

}

// updateArticle handles PUT /articles/:id - update article (protected)
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	// TODO: Get article ID from URL parameter
	// TODO: Parse JSON request body
	// TODO: Find and update article
	// TODO: Return updated article
	id := c.Param("id")
	artID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
		})
		return
	}

	var article dto.UpdateArticleRequest
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, dto.APIResponse{
			Success: true,
			Message: "json parse err",
			Error:   err.Error(),
		})
		return
	}

	art, err := h.articleService.UpdateArticle(artID, article)

	if err != nil {
		c.JSON(500, dto.APIResponse{
			Success: false,
			Message: "update fail",
			Error:   err.Error(),
		})
	}

	c.JSON(200, dto.APIResponse{
		Success: true,
		Data:    art,
	})
}
