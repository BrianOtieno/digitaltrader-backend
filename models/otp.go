package models

import (
	"time"

	"digitaltrader/db"
)

type OTP struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OTP        string    `json:"otp" gorm:"type:varchar(255)"`
	Email      string    `json:"email" gorm:"type:varchar(255)"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateOTP auto-migrates the OTP model
func MigrateOTP() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&OTP{})
}

// CreateOTP creates a new OTP record
func CreateOTP(otp *OTP) error {
	db := db.InitDB()
	return db.Create(otp).Error
}

// GetOTP retrieves an OTP record by ID
func GetOTP(id uint) (*OTP, error) {
	db := db.InitDB()
	var otp OTP
	if err := db.First(&otp, id).Error; err != nil {
		return nil, err
	}
	return &otp, nil
}

// UpdateOTP updates an existing OTP record
func UpdateOTP(otp *OTP, id uint) error {
	db := db.InitDB()
	return db.Model(&OTP{}).Where("id = ?", id).Updates(otp).Error
}

// DeleteOTP deletes an OTP record by ID
func DeleteOTP(id uint) error {
	db := db.InitDB()
	return db.Delete(&OTP{}, id).Error
}

// GetAllOTPs retrieves all OTP records
func GetAllOTPs() ([]OTP, error) {
	db := db.InitDB()
	var otps []OTP
	if err := db.Find(&otps).Error; err != nil {
		return nil, err
	}
	return otps, nil
}
