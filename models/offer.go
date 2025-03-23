package models

import (
	"time"

	"digitaltrader/db"
)

type Offer struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         *string    `json:"name" gorm:"type:text;null"`
	Off          *float64   `json:"off" gorm:"type:decimal(10,2);null"`
	Type         *string    `json:"type" gorm:"type:varchar(255);null"`
	Upto         *float64   `json:"upto" gorm:"type:decimal(10,2);null"`
	Min          *float64   `json:"min" gorm:"type:decimal(10,2);null"`
	From         *time.Time `json:"from" gorm:"type:date;null"`
	To           *time.Time `json:"to" gorm:"type:date;null"`
	DateTime     *string    `json:"date_time" gorm:"type:varchar(255);null"`
	Descriptions *string    `json:"descriptions" gorm:"type:text;null"`
	Image        *string    `json:"image" gorm:"type:text;null"`
	Manage       int8       `json:"manage" gorm:"type:tinyint;default:0"` // 0 = admin, 1 = store
	StoreID      *int       `json:"store_id" gorm:"type:integer;null"`
	Status       int8       `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField   *string    `json:"extra_field" gorm:"type:text;null"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	// Relationship
	Store Store `gorm:"foreignKey:StoreID;references:ID;constraint:OnDelete:SET NULL"`
}

// MigrateOffer auto-migrates the Offer model
func MigrateOffer() {
	db := db.InitDB()
	// Migrate Store first due to foreign key dependency
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Store{}, &Offer{})
}

// CreateOffer creates a new offer
func CreateOffer(offer *Offer) error {
	db := db.InitDB()
	return db.Create(offer).Error
}

// GetOffer retrieves an offer by ID
func GetOffer(id uint) (*Offer, error) {
	db := db.InitDB()
	var offer Offer
	if err := db.Preload("Store").First(&offer, id).Error; err != nil {
		return nil, err
	}
	return &offer, nil
}

// UpdateOffer updates an existing offer
func UpdateOffer(offer *Offer, id uint) error {
	db := db.InitDB()
	return db.Model(&Offer{}).Where("id = ?", id).Updates(offer).Error
}

// DeleteOffer deletes an offer by ID
func DeleteOffer(id uint) error {
	db := db.InitDB()
	return db.Delete(&Offer{}, id).Error
}

// GetAllOffers retrieves all offers
func GetAllOffers() ([]Offer, error) {
	db := db.InitDB()
	var offers []Offer
	if err := db.Preload("Store").Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}
