package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUsersPosts(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	c.JSON(http.StatusOK, []Post{})
}
