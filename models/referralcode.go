package models

import (
	"time"

	"digitaltrader/db"
)

type ReferralCode struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UID        int       `json:"uid" gorm:"type:integer"`
	Code       string    `json:"code" gorm:"type:text"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relationship (assuming a User model exists)
	User User `gorm:"foreignKey:UID;references:ID;constraint:OnDelete:RESTRICT"`
}

// MigrateReferralCode auto-migrates the ReferralCode model
func MigrateReferralCode() {
	db := db.InitDB()
	// Migrate User first due to foreign key dependency
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &ReferralCode{})
}

// CreateReferralCode creates a new referral code record
func CreateReferralCode(referralCode *ReferralCode) error {
	db := db.InitDB()
	return db.Create(referralCode).Error
}

// GetReferralCode retrieves a referral code record by ID
func GetReferralCode(id uint) (*ReferralCode, error) {
	db := db.InitDB()
	var referralCode ReferralCode
	if err := db.Preload("User").First(&referralCode, id).Error; err != nil {
		return nil, err
	}
	return &referralCode, nil
}

// UpdateReferralCode updates an existing referral code record
func UpdateReferralCode(referralCode *ReferralCode, id uint) error {
	db := db.InitDB()
	return db.Model(&ReferralCode{}).Where("id = ?", id).Updates(referralCode).Error
}

// DeleteReferralCode deletes a referral code record by ID
func DeleteReferralCode(id uint) error {
	db := db.InitDB()
	return db.Delete(&ReferralCode{}, id).Error
}

// GetAllReferralCodes retrieves all referral code records
func GetAllReferralCodes() ([]ReferralCode, error) {
	db := db.InitDB()
	var referralCodes []ReferralCode
	if err := db.Preload("User").Find(&referralCodes).Error; err != nil {
		return nil, err
	}
	return referralCodes, nil
}
