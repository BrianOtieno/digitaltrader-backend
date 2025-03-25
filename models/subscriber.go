package models

import (
	"time"

	"digitaltrader/db"
)

type Subscriber struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Email      string    `json:"email" gorm:"type:varchar(255)"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Timestamp  *string   `json:"timestamp" gorm:"type:varchar(255);null"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateSubscriber auto-migrates the Subscriber model
func MigrateSubscriber() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Subscriber{})
}

// CreateSubscriber creates a new subscriber record
func CreateSubscriber(subscriber *Subscriber) error {
	db := db.InitDB()
	return db.Create(subscriber).Error
}

// GetSubscriber retrieves a subscriber record by ID
func GetSubscriber(id uint) (*Subscriber, error) {
	db := db.InitDB()
	var subscriber Subscriber
	if err := db.First(&subscriber, id).Error; err != nil {
		return nil, err
	}
	return &subscriber, nil
}

// UpdateSubscriber updates an existing subscriber record
func UpdateSubscriber(subscriber *Subscriber, id uint) error {
	db := db.InitDB()
	return db.Model(&Subscriber{}).Where("id = ?", id).Updates(subscriber).Error
}

// DeleteSubscriber deletes a subscriber record by ID
func DeleteSubscriber(id uint) error {
	db := db.InitDB()
	return db.Delete(&Subscriber{}, id).Error
}

// GetAllSubscribers retrieves all subscriber records
func GetAllSubscribers() ([]Subscriber, error) {
	db := db.InitDB()
	var subscribers []Subscriber
	if err := db.Find(&subscribers).Error; err != nil {
		return nil, err
	}
	return subscribers, nil
}
