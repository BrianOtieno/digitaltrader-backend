package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStoreRequestController handles creating a new store request record
func CreateStoreRequestController(c *gin.Context) {
	var storeRequest models.StoreRequest
	if err := c.ShouldBindJSON(&storeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateStoreRequest(&storeRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, storeRequest)
}

// GetStoreRequestController retrieves a store request record by ID
func GetStoreRequestController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	storeRequest, err := models.GetStoreRequest(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store request not found"})
		return
	}
	c.JSON(http.StatusOK, storeRequest)
}

// UpdateStoreRequestController updates an existing store request record
func UpdateStoreRequestController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var storeRequest models.StoreRequest
	if err := c.ShouldBindJSON(&storeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateStoreRequest(&storeRequest, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, storeRequest)
}

// DeleteStoreRequestController deletes a store request record by ID
func DeleteStoreRequestController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteStoreRequest(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Store request deleted successfully"})
}

// GetAllStoreRequestsController retrieves all store request records
func GetAllStoreRequestsController(c *gin.Context) {
	storeRequests, err := models.GetAllStoreRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, storeRequests)
}
