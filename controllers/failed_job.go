package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateFailedJobController handles creating a new failed job
func CreateFailedJobController(c *gin.Context) {
	var failedJob models.FailedJob
	if err := c.ShouldBindJSON(&failedJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateFailedJob(&failedJob); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, failedJob)
}

// GetFailedJobController retrieves a failed job by ID
func GetFailedJobController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	failedJob, err := models.GetFailedJob(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed job not found"})
		return
	}
	c.JSON(http.StatusOK, failedJob)
}

// UpdateFailedJobController updates an existing failed job
func UpdateFailedJobController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var failedJob models.FailedJob
	if err := c.ShouldBindJSON(&failedJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateFailedJob(&failedJob, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, failedJob)
}

// DeleteFailedJobController deletes a failed job by ID
func DeleteFailedJobController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteFailedJob(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Failed job deleted successfully"})
}

// GetAllFailedJobsController retrieves all failed jobs
func GetAllFailedJobsController(c *gin.Context) {
	failedJobs, err := models.GetAllFailedJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, failedJobs)
}
