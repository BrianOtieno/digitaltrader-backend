package models

import (
	"time"

	"digitaltrader/db"
)

type Contact struct {
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       *string    `json:"name" gorm:"type:varchar(255);null"`
	Email      *string    `json:"email" gorm:"type:varchar(255);null"`
	Message    *string    `json:"message" gorm:"type:text;null"`
	Date       *time.Time `json:"date" gorm:"type:date;null"`
	ExtraField *string    `json:"extra_field" gorm:"type:text;null"`
	Status     int8       `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// MigrateContact auto-migrates the Contact model
func MigrateContact() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Contact{})
}

// CreateContact creates a new contact
func CreateContact(contact *Contact) error {
	db := db.InitDB()
	return db.Create(contact).Error
}

// GetContact retrieves a contact by ID
func GetContact(id uint) (*Contact, error) {
	db := db.InitDB()
	var contact Contact
	if err := db.First(&contact, id).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

// UpdateContact updates an existing contact
func UpdateContact(contact *Contact, id uint) error {
	db := db.InitDB()
	return db.Model(&Contact{}).Where("id = ?", id).Updates(contact).Error
}

// DeleteContact deletes a contact by ID
func DeleteContact(id uint) error {
	db := db.InitDB()
	return db.Delete(&Contact{}, id).Error
}

// GetAllContacts retrieves all contacts
func GetAllContacts() ([]Contact, error) {
	db := db.InitDB()
	var contacts []Contact
	if err := db.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}
