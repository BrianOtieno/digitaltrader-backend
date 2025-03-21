package models

import (
	"time"

	"digitaltrader/db"
)

type Popup struct {
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Shown      *int8      `json:"shown" gorm:"type:tinyint;null"`
	Message    *string    `json:"message" gorm:"type:text;null"`
	DateTime   *time.Time `json:"date_time" gorm:"type:datetime;null"`
	Status     int8       `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField *string    `json:"extra_field" gorm:"type:text;null"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// MigratePopup auto-migrates the Popup model
func MigratePopup() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Popup{})
}

// CreatePopup creates a new popup
func CreatePopup(popup *Popup) error {
	db := db.InitDB()
	return db.Create(popup).Error
}

// GetPopup retrieves a popup by ID
func GetPopup(id uint) (*Popup, error) {
	db := db.InitDB()
	var popup Popup
	if err := db.First(&popup, id).Error; err != nil {
		return nil, err
	}
	return &popup, nil
}

// UpdatePopup updates an existing popup
func UpdatePopup(popup *Popup, id uint) error {
	db := db.InitDB()
	return db.Model(&Popup{}).Where("id = ?", id).Updates(popup).Error
}

// DeletePopup deletes a popup by ID
func DeletePopup(id uint) error {
	db := db.InitDB()
	return db.Delete(&Popup{}, id).Error
}

// GetAllPopups retrieves all popups
func GetAllPopups() ([]Popup, error) {
	db := db.InitDB()
	var popups []Popup
	if err := db.Find(&popups).Error; err != nil {
		return nil, err
	}
	return popups, nil
}
