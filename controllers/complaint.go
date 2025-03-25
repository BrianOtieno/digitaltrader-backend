package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComplaintController handles creating a new complaint record
func CreateComplaintController(c *gin.Context) {
	var complaint models.Complaint
	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateComplaint(&complaint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, complaint)
}

// GetComplaintController retrieves a complaint record by ID
func GetComplaintController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	complaint, err := models.GetComplaint(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}
	c.JSON(http.StatusOK, complaint)
}

// UpdateComplaintController updates an existing complaint record
func UpdateComplaintController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var complaint models.Complaint
	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateComplaint(&complaint, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, complaint)
}

// DeleteComplaintController deletes a complaint record by ID
func DeleteComplaintController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteComplaint(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Complaint deleted successfully"})
}

// GetAllComplaintsController retrieves all complaint records
func GetAllComplaintsController(c *gin.Context) {
	complaints, err := models.GetAllComplaints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, complaints)
}
