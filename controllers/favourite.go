package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateFavouriteController handles creating a new favourite
func CreateFavouriteController(c *gin.Context) {
	var favourite models.Favourite
	if err := c.ShouldBindJSON(&favourite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateFavourite(&favourite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, favourite)
}

// GetFavouriteController retrieves a favourite by ID
func GetFavouriteController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	favourite, err := models.GetFavourite(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favourite not found"})
		return
	}
	c.JSON(http.StatusOK, favourite)
}

// UpdateFavouriteController updates an existing favourite
func UpdateFavouriteController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var favourite models.Favourite
	if err := c.ShouldBindJSON(&favourite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateFavourite(&favourite, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favourite)
}

// DeleteFavouriteController deletes a favourite by ID
func DeleteFavouriteController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteFavourite(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Favourite deleted successfully"})
}

// GetAllFavouritesController retrieves all favourites
func GetAllFavouritesController(c *gin.Context) {
	favourites, err := models.GetAllFavourites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favourites)
}
