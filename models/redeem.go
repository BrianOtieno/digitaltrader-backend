package models

import (
	"time"

	"digitaltrader/db"
)

type Redeem struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OwnerID    int       `json:"owner" gorm:"type:integer"`
	RedeemerID int       `json:"redeemer" gorm:"type:integer"`
	Code       string    `json:"code" gorm:"type:text"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relationships (assuming a User model exists)
	Owner    User `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:RESTRICT"`
	Redeemer User `gorm:"foreignKey:RedeemerID;references:ID;constraint:OnDelete:RESTRICT"`
}

// MigrateRedeem auto-migrates the Redeem model
func MigrateRedeem() {
	db := db.InitDB()
	// Migrate User first due to foreign key dependencies
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Redeem{})
}

// CreateRedeem creates a new redeem record
func CreateRedeem(redeem *Redeem) error {
	db := db.InitDB()
	return db.Create(redeem).Error
}

// GetRedeem retrieves a redeem record by ID
func GetRedeem(id uint) (*Redeem, error) {
	db := db.InitDB()
	var redeem Redeem
	if err := db.Preload("Owner").Preload("Redeemer").First(&redeem, id).Error; err != nil {
		return nil, err
	}
	return &redeem, nil
}

// UpdateRedeem updates an existing redeem record
func UpdateRedeem(redeem *Redeem, id uint) error {
	db := db.InitDB()
	return db.Model(&Redeem{}).Where("id = ?", id).Updates(redeem).Error
}

// DeleteRedeem deletes a redeem record by ID
func DeleteRedeem(id uint) error {
	db := db.InitDB()
	return db.Delete(&Redeem{}, id).Error
}

// GetAllRedeems retrieves all redeem records
func GetAllRedeems() ([]Redeem, error) {
	db := db.InitDB()
	var redeems []Redeem
	if err := db.Preload("Owner").Preload("Redeemer").Find(&redeems).Error; err != nil {
		return nil, err
	}
	return redeems, nil
}
