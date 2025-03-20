package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateFlushController handles creating a new flush
func CreateFlushController(c *gin.Context) {
	var flush models.Flush
	if err := c.ShouldBindJSON(&flush); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateFlush(&flush); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, flush)
}

// GetFlushController retrieves a flush by ID
func GetFlushController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	flush, err := models.GetFlush(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flush not found"})
		return
	}
	c.JSON(http.StatusOK, flush)
}

// UpdateFlushController updates an existing flush
func UpdateFlushController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var flush models.Flush
	if err := c.ShouldBindJSON(&flush); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateFlush(&flush, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flush)
}

// DeleteFlushController deletes a flush by ID
func DeleteFlushController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteFlush(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Flush deleted successfully"})
}

// GetAllFlushesController retrieves all flushes
func GetAllFlushesController(c *gin.Context) {
	flushes, err := models.GetAllFlushes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flushes)
}
