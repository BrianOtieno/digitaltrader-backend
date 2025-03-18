package models

import (
	"time"

	"digitaltrader/db"
)

type Address struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UID        int       `json:"uid" gorm:"type:integer"`
	Title      *string   `json:"title" gorm:"type:varchar(255);null"`
	Address    *string   `json:"address" gorm:"type:text;null"`
	House      *string   `json:"house" gorm:"type:text;null"`
	Landmark   *string   `json:"landmark" gorm:"type:text;null"`
	Pincode    *string   `json:"pincode" gorm:"type:varchar(255);null"`
	Lat        string    `json:"lat" gorm:"type:varchar(255)"`
	Lng        string    `json:"lng" gorm:"type:varchar(255)"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateAddress auto-migrates the Address model
func MigrateAddress() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Address{})
}

// CreateAddress creates a new address
func CreateAddress(address *Address) error {
	db := db.InitDB()
	return db.Create(address).Error
}

// GetAddress retrieves an address by ID
func GetAddress(id uint) (*Address, error) {
	db := db.InitDB()
	var address Address
	if err := db.First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

// UpdateAddress updates an existing address
func UpdateAddress(address *Address, id uint) error {
	db := db.InitDB()
	return db.Model(&Address{}).Where("id = ?", id).Updates(address).Error
}

// DeleteAddress deletes an address by ID
func DeleteAddress(id uint) error {
	db := db.InitDB()
	return db.Delete(&Address{}, id).Error
}

// GetAllAddresses retrieves all addresses
func GetAllAddresses() ([]Address, error) {
	db := db.InitDB()
	var addresses []Address
	if err := db.Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}
