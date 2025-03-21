package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAnalyticsController handles creating a new analytics record
func CreateAnalyticsController(c *gin.Context) {
	var analytics models.Analytics
	if err := c.ShouldBindJSON(&analytics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateAnalytics(&analytics); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, analytics)
}

// GetAnalyticsController retrieves an analytics record by ID
func GetAnalyticsController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	analytics, err := models.GetAnalytics(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Analytics record not found"})
		return
	}
	c.JSON(http.StatusOK, analytics)
}

// UpdateAnalyticsController updates an existing analytics record
func UpdateAnalyticsController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var analytics models.Analytics
	if err := c.ShouldBindJSON(&analytics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateAnalytics(&analytics, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, analytics)
}

// DeleteAnalyticsController deletes an analytics record by ID
func DeleteAnalyticsController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteAnalytics(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Analytics record deleted successfully"})
}

// GetAllAnalyticsController retrieves all analytics records
func GetAllAnalyticsController(c *gin.Context) {
	analytics, err := models.GetAllAnalytics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, analytics)
}
