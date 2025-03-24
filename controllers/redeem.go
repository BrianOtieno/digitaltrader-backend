package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRedeemController handles creating a new redeem record
func CreateRedeemController(c *gin.Context) {
	var redeem models.Redeem
	if err := c.ShouldBindJSON(&redeem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateRedeem(&redeem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, redeem)
}

// GetRedeemController retrieves a redeem record by ID
func GetRedeemController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	redeem, err := models.GetRedeem(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Redeem record not found"})
		return
	}
	c.JSON(http.StatusOK, redeem)
}

// UpdateRedeemController updates an existing redeem record
func UpdateRedeemController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var redeem models.Redeem
	if err := c.ShouldBindJSON(&redeem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateRedeem(&redeem, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, redeem)
}

// DeleteRedeemController deletes a redeem record by ID
func DeleteRedeemController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteRedeem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Redeem record deleted successfully"})
}

// GetAllRedeemsController retrieves all redeem records
func GetAllRedeemsController(c *gin.Context) {
	redeems, err := models.GetAllRedeems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, redeems)
}
