package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateReferralCodeController handles creating a new referral code record
func CreateReferralCodeController(c *gin.Context) {
	var referralCode models.ReferralCode
	if err := c.ShouldBindJSON(&referralCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateReferralCode(&referralCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, referralCode)
}

// GetReferralCodeController retrieves a referral code record by ID
func GetReferralCodeController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	referralCode, err := models.GetReferralCode(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Referral code not found"})
		return
	}
	c.JSON(http.StatusOK, referralCode)
}

// UpdateReferralCodeController updates an existing referral code record
func UpdateReferralCodeController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var referralCode models.ReferralCode
	if err := c.ShouldBindJSON(&referralCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateReferralCode(&referralCode, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, referralCode)
}

// DeleteReferralCodeController deletes a referral code record by ID
func DeleteReferralCodeController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteReferralCode(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Referral code deleted successfully"})
}

// GetAllReferralCodesController retrieves all referral code records
func GetAllReferralCodesController(c *gin.Context) {
	referralCodes, err := models.GetAllReferralCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, referralCodes)
}
