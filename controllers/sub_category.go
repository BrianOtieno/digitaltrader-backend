package controllers

import (
	"digitaltrader/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSubCategoryController handles creating a new sub-category
func CreateSubCategoryController(c *gin.Context) {
	var subCategory models.SubCategory
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateSubCategory(&subCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, subCategory)
}

// GetSubCategoryController retrieves a sub-category by ID
func GetSubCategoryController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	subCategory, err := models.GetSubCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub-category not found"})
		return
	}
	c.JSON(http.StatusOK, subCategory)
}

// UpdateSubCategoryController updates an existing sub-category
func UpdateSubCategoryController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var subCategory models.SubCategory
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateSubCategory(&subCategory, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subCategory)
}

// DeleteSubCategoryController deletes a sub-category by ID
func DeleteSubCategoryController(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteSubCategory(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sub-category deleted successfully"})
}

// GetAllSubCategoriesController retrieves all sub-categories
func GetAllSubCategoriesController(c *gin.Context) {
	subCategories, err := models.GetAllSubCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subCategories)
}
