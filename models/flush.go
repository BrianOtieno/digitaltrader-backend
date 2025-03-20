package models

import (
	"time"

	"digitaltrader/db"
)

type Flush struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Key        string    `json:"key" gorm:"type:text"`
	Value      string    `json:"value" gorm:"type:text"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateFlush auto-migrates the Flush model
func MigrateFlush() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Flush{})
}

// CreateFlush creates a new flush
func CreateFlush(flush *Flush) error {
	db := db.InitDB()
	return db.Create(flush).Error
}

// GetFlush retrieves a flush by ID
func GetFlush(id uint) (*Flush, error) {
	db := db.InitDB()
	var flush Flush
	if err := db.First(&flush, id).Error; err != nil {
		return nil, err
	}
	return &flush, nil
}

// UpdateFlush updates an existing flush
func UpdateFlush(flush *Flush, id uint) error {
	db := db.InitDB()
	return db.Model(&Flush{}).Where("id = ?", id).Updates(flush).Error
}

// DeleteFlush deletes a flush by ID
func DeleteFlush(id uint) error {
	db := db.InitDB()
	return db.Delete(&Flush{}, id).Error
}

// GetAllFlushes retrieves all flushes
func GetAllFlushes() ([]Flush, error) {
	db := db.InitDB()
	var flushes []Flush
	if err := db.Find(&flushes).Error; err != nil {
		return nil, err
	}
	return flushes, nil
}
