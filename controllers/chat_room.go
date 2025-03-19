package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateChatRoomController handles creating a new chat room
func CreateChatRoomController(c *gin.Context) {
	var chatRoom models.ChatRoom
	if err := c.ShouldBindJSON(&chatRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateChatRoom(&chatRoom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, chatRoom)
}

// GetChatRoomController retrieves a chat room by ID
func GetChatRoomController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	chatRoom, err := models.GetChatRoom(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chat room not found"})
		return
	}
	c.JSON(http.StatusOK, chatRoom)
}

// UpdateChatRoomController updates an existing chat room
func UpdateChatRoomController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var chatRoom models.ChatRoom
	if err := c.ShouldBindJSON(&chatRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateChatRoom(&chatRoom, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chatRoom)
}

// DeleteChatRoomController deletes a chat room by ID
func DeleteChatRoomController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteChatRoom(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Chat room deleted successfully"})
}

// GetAllChatRoomsController retrieves all chat rooms
func GetAllChatRoomsController(c *gin.Context) {
	chatRooms, err := models.GetAllChatRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chatRooms)
}
