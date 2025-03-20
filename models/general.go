package models

import (
	"time"

	"digitaltrader/db"
)

type General struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          *string   `json:"name" gorm:"type:varchar(255);null"`
	Mobile        *string   `json:"mobile" gorm:"type:varchar(255);null"`
	Email         *string   `json:"email" gorm:"type:varchar(255);null"`
	Address       *string   `json:"address" gorm:"type:varchar(255);null"`
	City          *string   `json:"city" gorm:"type:varchar(255);null"`
	State         *string   `json:"state" gorm:"type:varchar(255);null"`
	Zip           *string   `json:"zip" gorm:"type:varchar(255);null"`
	Country       *string   `json:"country" gorm:"type:varchar(255);null"`
	Min           *float64  `json:"min" gorm:"type:decimal(10,2);null"`
	Free          *float64  `json:"free" gorm:"type:decimal(10,2);null"`
	Tax           *float64  `json:"tax" gorm:"type:decimal(10,2);null"`
	Shipping      *string   `json:"shipping" gorm:"type:varchar(255);null"`
	ShippingPrice *float64  `json:"shippingPrice" gorm:"type:decimal(10,2);null"`
	Status        int8      `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField    *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// MigrateGeneral auto-migrates the General model
func MigrateGeneral() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&General{})
}

// CreateGeneral creates a new general record
func CreateGeneral(general *General) error {
	db := db.InitDB()
	return db.Create(general).Error
}

// GetGeneral retrieves a general record by ID
func GetGeneral(id uint) (*General, error) {
	db := db.InitDB()
	var general General
	if err := db.First(&general, id).Error; err != nil {
		return nil, err
	}
	return &general, nil
}

// UpdateGeneral updates an existing general record
func UpdateGeneral(general *General, id uint) error {
	db := db.InitDB()
	return db.Model(&General{}).Where("id = ?", id).Updates(general).Error
}

// DeleteGeneral deletes a general record by ID
func DeleteGeneral(id uint) error {
	db := db.InitDB()
	return db.Delete(&General{}, id).Error
}

// GetAllGenerals retrieves all general records
func GetAllGenerals() ([]General, error) {
	db := db.InitDB()
	var generals []General
	if err := db.Find(&generals).Error; err != nil {
		return nil, err
	}
	return generals, nil
}
