package models

import (
	"time"

	"digitaltrader/db"
)

type Category struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       *string   `json:"name" gorm:"type:varchar(255);null"`
	Cover      *string   `json:"cover" gorm:"type:varchar(255);null"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateCategory auto-migrates the Category model
func MigrateCategory() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Category{})
}

// CreateCategory creates a new category
func CreateCategory(category *Category) error {
	db := db.InitDB()
	return db.Create(category).Error
}

// GetCategory retrieves a category by ID
func GetCategory(id uint) (*Category, error) {
	db := db.InitDB()
	var category Category
	if err := db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// UpdateCategory updates an existing category
func UpdateCategory(category *Category, id uint) error {
	db := db.InitDB()
	return db.Model(&Category{}).Where("id = ?", id).Updates(category).Error
}

// DeleteCategory deletes a category by ID
func DeleteCategory(id uint) error {
	db := db.InitDB()
	return db.Delete(&Category{}, id).Error
}

// GetAllCategories retrieves all categories
func GetAllCategories() ([]Category, error) {
	db := db.InitDB()
	var categories []Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
