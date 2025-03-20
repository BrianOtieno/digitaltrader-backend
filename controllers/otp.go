package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOTPController handles creating a new OTP record
func CreateOTPController(c *gin.Context) {
	var otp models.OTP
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateOTP(&otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, otp)
}

// GetOTPController retrieves an OTP record by ID
func GetOTPController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	otp, err := models.GetOTP(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OTP not found"})
		return
	}
	c.JSON(http.StatusOK, otp)
}

// UpdateOTPController updates an existing OTP record
func UpdateOTPController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var otp models.OTP
	if err := c.ShouldBindJSON(&otp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateOTP(&otp, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, otp)
}

// DeleteOTPController deletes an OTP record by ID
func DeleteOTPController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteOTP(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP deleted successfully"})
}

// GetAllOTPsController retrieves all OTP records
func GetAllOTPsController(c *gin.Context) {
	otps, err := models.GetAllOTPs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, otps)
}
