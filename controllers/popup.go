package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePopupController handles creating a new popup
func CreatePopupController(c *gin.Context) {
	var popup models.Popup
	if err := c.ShouldBindJSON(&popup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreatePopup(&popup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, popup)
}

// GetPopupController retrieves a popup by ID
func GetPopupController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	popup, err := models.GetPopup(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Popup not found"})
		return
	}
	c.JSON(http.StatusOK, popup)
}

// UpdatePopupController updates an existing popup
func UpdatePopupController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var popup models.Popup
	if err := c.ShouldBindJSON(&popup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdatePopup(&popup, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, popup)
}

// DeletePopupController deletes a popup by ID
func DeletePopupController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeletePopup(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Popup deleted successfully"})
}

// GetAllPopupsController retrieves all popups
func GetAllPopupsController(c *gin.Context) {
	popups, err := models.GetAllPopups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, popups)
}
