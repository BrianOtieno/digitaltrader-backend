package models

import (
	"time"

	"digitaltrader/db"
)

type City struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Lat        string    `json:"lat" gorm:"type:varchar(255)"`
	Lng        string    `json:"lng" gorm:"type:varchar(255)"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateCity auto-migrates the City model
func MigrateCity() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&City{})
}

// CreateCity creates a new city
func CreateCity(city *City) error {
	db := db.InitDB()
	return db.Create(city).Error
}

// GetCity retrieves a city by ID
func GetCity(id uint) (*City, error) {
	db := db.InitDB()
	var city City
	if err := db.First(&city, id).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

// UpdateCity updates an existing city
func UpdateCity(city *City, id uint) error {
	db := db.InitDB()
	return db.Model(&City{}).Where("id = ?", id).Updates(city).Error
}

// DeleteCity deletes a city by ID
func DeleteCity(id uint) error {
	db := db.InitDB()
	return db.Delete(&City{}, id).Error
}

// GetAllCities retrieves all cities
func GetAllCities() ([]City, error) {
	db := db.InitDB()
	var cities []City
	if err := db.Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}
