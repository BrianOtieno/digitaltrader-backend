package models

import (
	"time"

	"digitaltrader/db"
)

type SubCategory struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Cover      string    `json:"cover" gorm:"type:varchar(255)"`
	CateID     int       `json:"cate_id" gorm:"type:integer"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Status     int8      `json:"status" gorm:"type:tinyint"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relationship
	Category Category `gorm:"foreignKey:CateID;references:ID;constraint:OnDelete:RESTRICT"`
}

// MigrateSubCategory auto-migrates the SubCategory model
func MigrateSubCategory() {
	db := db.InitDB()
	// Migrate Category first due to foreign key dependency
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Category{}, &SubCategory{})
}

// CreateSubCategory creates a new sub-category
func CreateSubCategory(subCategory *SubCategory) error {
	db := db.InitDB()
	return db.Create(subCategory).Error
}

// GetSubCategory retrieves a sub-category by ID
func GetSubCategory(id uint) (*SubCategory, error) {
	db := db.InitDB()
	var subCategory SubCategory
	if err := db.Preload("Category").First(&subCategory, id).Error; err != nil {
		return nil, err
	}
	return &subCategory, nil
}

// UpdateSubCategory updates an existing sub-category
func UpdateSubCategory(subCategory *SubCategory, id uint) error {
	db := db.InitDB()
	return db.Model(&SubCategory{}).Where("id = ?", id).Updates(subCategory).Error
}

// DeleteSubCategory deletes a sub-category by ID
func DeleteSubCategory(id uint) error {
	db := db.InitDB()
	return db.Delete(&SubCategory{}, id).Error
}

// GetAllSubCategories retrieves all sub-categories
func GetAllSubCategories() ([]SubCategory, error) {
	db := db.InitDB()
	var subCategories []SubCategory
	if err := db.Preload("Category").Find(&subCategories).Error; err != nil {
		return nil, err
	}
	return subCategories, nil
}
