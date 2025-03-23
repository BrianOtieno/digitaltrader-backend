package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateManageController handles creating a new manage record
func CreateManageController(c *gin.Context) {
	var manage models.Manage
	if err := c.ShouldBindJSON(&manage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateManage(&manage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, manage)
}

// GetManageController retrieves a manage record by ID
func GetManageController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	manage, err := models.GetManage(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manage record not found"})
		return
	}
	c.JSON(http.StatusOK, manage)
}

// UpdateManageController updates an existing manage record
func UpdateManageController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var manage models.Manage
	if err := c.ShouldBindJSON(&manage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateManage(&manage, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, manage)
}

// DeleteManageController deletes a manage record by ID
func DeleteManageController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteManage(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Manage record deleted successfully"})
}

// GetAllManagesController retrieves all manage records
func GetAllManagesController(c *gin.Context) {
	manages, err := models.GetAllManages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, manages)
}
