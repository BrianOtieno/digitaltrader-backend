package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDriverRequestController handles creating a new driver request record
func CreateDriverRequestController(c *gin.Context) {
	var driverRequest models.DriverRequest
	if err := c.ShouldBindJSON(&driverRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateDriverRequest(&driverRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, driverRequest)
}

// GetDriverRequestController retrieves a driver request record by ID
func GetDriverRequestController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	driverRequest, err := models.GetDriverRequest(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver request not found"})
		return
	}
	c.JSON(http.StatusOK, driverRequest)
}

// UpdateDriverRequestController updates an existing driver request record
func UpdateDriverRequestController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var driverRequest models.DriverRequest
	if err := c.ShouldBindJSON(&driverRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateDriverRequest(&driverRequest, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driverRequest)
}

// DeleteDriverRequestController deletes a driver request record by ID
func DeleteDriverRequestController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteDriverRequest(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Driver request deleted successfully"})
}

// GetAllDriverRequestsController retrieves all driver request records
func GetAllDriverRequestsController(c *gin.Context) {
	driverRequests, err := models.GetAllDriverRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driverRequests)
}
