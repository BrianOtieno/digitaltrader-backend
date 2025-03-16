// Controller functions
package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTransactionController handles creating a new transaction
func CreateTransactionController(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaction)
}

// GetTransactionController retrieves a transaction by ID
func GetTransactionController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var transaction models.Transaction
	if err := models.GetTransaction(&transaction, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// UpdateTransactionController updates an existing transaction
func UpdateTransactionController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.UpdateTransaction(&transaction, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// DeleteTransactionController deletes a transaction by ID
func DeleteTransactionController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := models.DeleteTransaction(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

// GetAllTransactionsController retrieves all transactions
func GetAllTransactionsController(c *gin.Context) {
	var transactions []models.Transaction
	if err := models.GetAllTransactions(&transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
