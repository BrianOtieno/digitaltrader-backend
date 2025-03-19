package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCityController handles creating a new city
func CreateCityController(c *gin.Context) {
	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateCity(&city); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, city)
}

// GetCityController retrieves a city by ID
func GetCityController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	city, err := models.GetCity(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
		return
	}
	c.JSON(http.StatusOK, city)
}

// UpdateCityController updates an existing city
func UpdateCityController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var city models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateCity(&city, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, city)
}

// DeleteCityController deletes a city by ID
func DeleteCityController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteCity(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "City deleted successfully"})
}

// GetAllCitiesController retrieves all cities
func GetAllCitiesController(c *gin.Context) {
	cities, err := models.GetAllCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cities)
}
