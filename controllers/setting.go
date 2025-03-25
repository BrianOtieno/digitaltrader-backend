package controllers

import (
	"digitaltrader/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSettingController handles creating or updating the settings record
func CreateSettingController(c *gin.Context) {
	var setting models.Setting
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateSetting(&setting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, setting)
}

// GetSettingController retrieves the settings record
func GetSettingController(c *gin.Context) {
	setting, err := models.GetSetting()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Settings not found"})
		return
	}
	c.JSON(http.StatusOK, setting)
}

// UpdateSettingController updates the settings record
func UpdateSettingController(c *gin.Context) {
	var setting models.Setting
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.UpdateSetting(&setting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, setting)
}

// DeleteSettingController deletes the settings record
func DeleteSettingController(c *gin.Context) {
	if err := models.DeleteSetting(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Settings deleted successfully"})
}

// GetAllSettingsController retrieves all settings (typically one record)
func GetAllSettingsController(c *gin.Context) {
	settings, err := models.GetAllSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}
