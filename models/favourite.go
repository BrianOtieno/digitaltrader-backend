package models

import (
	"time"

	"digitaltrader/db"
)

type Favourite struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UID        int       `json:"uid" gorm:"type:integer"`
	IDs        *string   `json:"ids" gorm:"type:text;null"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateFavourite auto-migrates the Favourite model
func MigrateFavourite() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Favourite{})
}

// CreateFavourite creates a new favourite
func CreateFavourite(favourite *Favourite) error {
	db := db.InitDB()
	return db.Create(favourite).Error
}

// GetFavourite retrieves a favourite by ID
func GetFavourite(id uint) (*Favourite, error) {
	db := db.InitDB()
	var favourite Favourite
	if err := db.First(&favourite, id).Error; err != nil {
		return nil, err
	}
	return &favourite, nil
}

// UpdateFavourite updates an existing favourite
func UpdateFavourite(favourite *Favourite, id uint) error {
	db := db.InitDB()
	return db.Model(&Favourite{}).Where("id = ?", id).Updates(favourite).Error
}

// DeleteFavourite deletes a favourite by ID
func DeleteFavourite(id uint) error {
	db := db.InitDB()
	return db.Delete(&Favourite{}, id).Error
}

// GetAllFavourites retrieves all favourites
func GetAllFavourites() ([]Favourite, error) {
	db := db.InitDB()
	var favourites []Favourite
	if err := db.Find(&favourites).Error; err != nil {
		return nil, err
	}
	return favourites, nil
}
