package models

import (
	"time"

	"digitaltrader/db"
)

type Analytics struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Analytics *string   `json:"analytics" gorm:"type:text;null"`
	IP        *string   `json:"ip" gorm:"type:text;null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MigrateAnalytics auto-migrates the Analytics model
func MigrateAnalytics() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Analytics{})
}

// CreateAnalytics creates a new analytics record
func CreateAnalytics(analytics *Analytics) error {
	db := db.InitDB()
	return db.Create(analytics).Error
}

// GetAnalytics retrieves an analytics record by ID
func GetAnalytics(id uint) (*Analytics, error) {
	db := db.InitDB()
	var analytics Analytics
	if err := db.First(&analytics, id).Error; err != nil {
		return nil, err
	}
	return &analytics, nil
}

// UpdateAnalytics updates an existing analytics record
func UpdateAnalytics(analytics *Analytics, id uint) error {
	db := db.InitDB()
	return db.Model(&Analytics{}).Where("id = ?", id).Updates(analytics).Error
}

// DeleteAnalytics deletes an analytics record by ID
func DeleteAnalytics(id uint) error {
	db := db.InitDB()
	return db.Delete(&Analytics{}, id).Error
}

// GetAllAnalytics retrieves all analytics records
func GetAllAnalytics() ([]Analytics, error) {
	db := db.InitDB()
	var analytics []Analytics
	if err := db.Find(&analytics).Error; err != nil {
		return nil, err
	}
	return analytics, nil
}
