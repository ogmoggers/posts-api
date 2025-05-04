package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-network-api/pkg/middleware"
)

type Post struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (h *Handler) createPost(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var postReq PostRequest
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := Post{
		Title:   postReq.Title,
		Content: postReq.Content,
		UserID:  userID,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "post": post})
}

func (h *Handler) updatePost(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	postID := c.Param("post_id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	var postReq PostRequest
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := Post{
		ID:      postID,
		Title:   postReq.Title,
		Content: postReq.Content,
		UserID:  userID,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

func (h *Handler) getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, []Post{})
}

func (h *Handler) getPostById(c *gin.Context) {
	postID := c.Param("post_id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	c.JSON(http.StatusOK, Post{})
}

func (h *Handler) deletePost(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	postID := c.Param("post_id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully by user " + userID})
}
