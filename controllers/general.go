package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateGeneralController handles creating a new general record
func CreateGeneralController(c *gin.Context) {
	var general models.General
	if err := c.ShouldBindJSON(&general); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateGeneral(&general); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, general)
}

// GetGeneralController retrieves a general record by ID
func GetGeneralController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	general, err := models.GetGeneral(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "General record not found"})
		return
	}
	c.JSON(http.StatusOK, general)
}

// UpdateGeneralController updates an existing general record
func UpdateGeneralController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var general models.General
	if err := c.ShouldBindJSON(&general); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateGeneral(&general, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, general)
}

// DeleteGeneralController deletes a general record by ID
func DeleteGeneralController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteGeneral(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "General record deleted successfully"})
}

// GetAllGeneralsController retrieves all general records
func GetAllGeneralsController(c *gin.Context) {
	generals, err := models.GetAllGenerals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, generals)
}
