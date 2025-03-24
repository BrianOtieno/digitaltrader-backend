package models

import (
	"time"

	"digitaltrader/db"
)

type Referral struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Amount      float64   `json:"amount" gorm:"type:decimal(10,2)"`
	Title       string    `json:"title" gorm:"type:text"`
	Message     string    `json:"message" gorm:"type:text"`
	Limit       int       `json:"limit" gorm:"type:integer"`
	WhoReceived int8      `json:"who_received" gorm:"type:tinyint"`
	Status      int8      `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField  *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `gorm:"foreignKey:WhoReceived;references:ID;constraint:OnDelete:RESTRICT"`
}

// MigrateReferral auto-migrates the Referral model
func MigrateReferral() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Referral{})
}

// CreateReferral creates a new referral record
func CreateReferral(referral *Referral) error {
	db := db.InitDB()
	return db.Create(referral).Error
}

// GetReferral retrieves a referral record by ID
func GetReferral(id uint) (*Referral, error) {
	db := db.InitDB()
	var referral Referral
	if err := db.First(&referral, id).Error; err != nil {
		return nil, err
	}
	return &referral, nil
}

// UpdateReferral updates an existing referral record
func UpdateReferral(referral *Referral, id uint) error {
	db := db.InitDB()
	return db.Model(&Referral{}).Where("id = ?", id).Updates(referral).Error
}

// DeleteReferral deletes a referral record by ID
func DeleteReferral(id uint) error {
	db := db.InitDB()
	return db.Delete(&Referral{}, id).Error
}

// GetAllReferrals retrieves all referral records
func GetAllReferrals() ([]Referral, error) {
	db := db.InitDB()
	var referrals []Referral
	if err := db.Find(&referrals).Error; err != nil {
		return nil, err
	}
	return referrals, nil
}
