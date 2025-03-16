package controllers

import (
	"digitaltrader/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransferHandler handles creating a new transfer
func CreateTransferHandler(c *gin.Context) {
	var transfer models.Transfer
	if err := c.ShouldBindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateTransfer(&transfer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transfer)
}

// GetTransferByIDHandler handles retrieving a transfer by ID
func GetTransferByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var transfer models.Transfer

	if err := models.GetTransferByID(&transfer, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transfer not found"})
		return
	}

	c.JSON(http.StatusOK, transfer)
}

// GetAllTransfersHandler handles retrieving all transfers
func GetAllTransfersHandler(c *gin.Context) {
	var transfers []models.Transfer

	if err := models.GetAllTransfers(&transfers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transfers)
}

// UpdateTransferHandler handles updating a transfer by ID
func UpdateTransferHandler(c *gin.Context) {
	id := c.Param("id")
	var transfer models.Transfer

	if err := c.ShouldBindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateTransfer(&transfer, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transfer)
}

// DeleteTransferHandler handles deleting a transfer by ID
func DeleteTransferHandler(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteTransfer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Transfer deleted"})
}
