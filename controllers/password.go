package controllers

import (
	"net/http"

	"digitaltrader/models"

	"github.com/gin-gonic/gin"
)

// CreatePasswordResetTokenController handles creating a new password reset token
func CreatePasswordResetTokenController(c *gin.Context) {
	var token models.PasswordResetToken
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreatePasswordResetToken(&token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create password reset token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Password reset token created successfully"})
}

// GetPasswordResetTokenController handles retrieving a password reset token by email
func GetPasswordResetTokenController(c *gin.Context) {
	email := c.Param("email")
	var token models.PasswordResetToken

	if err := models.GetPasswordResetToken(&token, email); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Password reset token not found"})
		return
	}

	c.JSON(http.StatusOK, token)
}

// UpdatePasswordResetTokenController handles updating an existing password reset token
func UpdatePasswordResetTokenController(c *gin.Context) {
	email := c.Param("email")
	var token models.PasswordResetToken
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdatePasswordResetToken(&token, email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password reset token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset token updated successfully"})
}

// DeletePasswordResetTokenController handles deleting a password reset token by email
func DeletePasswordResetTokenController(c *gin.Context) {
	email := c.Param("email")

	if err := models.DeletePasswordResetToken(email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete password reset token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset token deleted successfully"})
}

// GetAllPasswordResetTokensController retrieves all password reset tokens
func GetAllPasswordResetTokensController(c *gin.Context) {
	var tokens []models.PasswordResetToken

	if err := models.GetAllPasswordResetTokens(&tokens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tokens"})
		return
	}

	c.JSON(http.StatusOK, tokens)
}
