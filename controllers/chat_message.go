package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateChatMessageController handles creating a new chat message
func CreateChatMessageController(c *gin.Context) {
	var chatMessage models.ChatMessage
	if err := c.ShouldBindJSON(&chatMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateChatMessage(&chatMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, chatMessage)
}

// GetChatMessageController retrieves a chat message by ID
func GetChatMessageController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	chatMessage, err := models.GetChatMessage(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chat message not found"})
		return
	}
	c.JSON(http.StatusOK, chatMessage)
}

// UpdateChatMessageController updates an existing chat message
func UpdateChatMessageController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var chatMessage models.ChatMessage
	if err := c.ShouldBindJSON(&chatMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateChatMessage(&chatMessage, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chatMessage)
}

// DeleteChatMessageController deletes a chat message by ID
func DeleteChatMessageController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteChatMessage(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Chat message deleted successfully"})
}

// GetAllChatMessagesController retrieves all chat messages
func GetAllChatMessagesController(c *gin.Context) {
	chatMessages, err := models.GetAllChatMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chatMessages)
}
